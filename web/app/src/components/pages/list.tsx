//src/components/pages/list.tsx
import React, {useEffect, useState} from 'react';
import {Table, Row, Col, Button, Typography} from 'antd';
import {useNavigate} from 'react-router';
import axios from 'axios';

const {Title} = Typography;


const List = () => {
  const history = useNavigate();
  const [allData, setAllData] = useState([]);
  const [selectedUser, setSelectedUser] = useState(false);
  const [userData, setUserData] = useState<UserData | null>(null);
  //Retrieve shadow realm from localstorage.
  const [lockedAccounts, setLockedAccounts] = useState<string[]>(JSON.parse(localStorage.getItem('shadowRealm') || '[]'));
  const [managerCheck, setLoggedInUserRole] = useState<string | null>(null);

  useEffect(() => {
      axios.get(`http://localhost:8080/employees`).then(res => {
        setAllData(res.data);
        setLoggedInUserRole(sessionStorage.getItem("UserRole"));
      });
    }, []);

  interface UserData {
    key: any;
    Username: any;
    FirstName: any;
    LastName: any;
    Role: any;
  }

  const columns = [
    {
      title: 'Username',
      dataIndex: 'Username',
    },
    {
      title: 'First Name',
      dataIndex: 'FirstName'
    },
    {
      title: 'Last Name',
      dataIndex: 'LastName'
    },
    {
      title: 'Role',
      dataIndex: 'Role'
    },
  ];

  const data: { key: any; Username: any; FirstName: any; LastName: any; Role: any; }[] = [];
    allData.map((user: any) => {
        data.push({
        key: user.ID,
        Username: user.Username,
        FirstName: user.FirstName,
        LastName:  user.LastName,
        Role:      user.Role,
      })
      return data;
    });
  console.log(data)

  const handleAddClick = () => {
    history('/form')
    }
  
  const handleDelClick = () => {
    if (!userData) {
      console.error('No user selected for deletion');
      return;
    }

  const handleModClick = () => {
    sessionStorage.setItem('firstNameEdit', userData.FirstName);
    sessionStorage.setItem('surnameEdit', userData.LastName);
    sessionStorage.setItem('roleEdit', userData.Role);
    history('/form')    
  }

    const deleteData = {
      id: userData.key,
      firstName: userData.FirstName
    }

    axios.delete(`http://localhost:8080/deleteEmployee`, {
      data: deleteData,
    })
    .then(res => {
      axios.get('http://localhost:8080/employees').then((res) => {
        setAllData(res.data);
      });
    })
    .catch((error) => {
      console.error('Error deleting employee:', error);
    });
    }
  

  const handleRowClick = (record : UserData) => {
    setSelectedUser(record.Username) // using horrible JS true value interpretation
    setUserData(record);
  }

  const handleUnlockClick = () => {
    if (!userData) {
      console.error('No user selected');
      return;
    }
    //Remove accounts from locked list.
    const updatedLockedAccounts = lockedAccounts.filter(account => account !== userData.Username);
    setLockedAccounts(updatedLockedAccounts);
    localStorage.setItem('shadowRealm', JSON.stringify(updatedLockedAccounts));
  };
    
  return (
      <div>
          <Row gutter={[40, 0]}>
            <Col span={10}>
              <Title level={2}>
              User List
              </Title>
              </Col>
            <Col span={15}>
              {selectedUser && <Title level={3}>Selected User: {selectedUser}</Title>}
            </Col>
            <Col span={2}>
            {managerCheck === "Manager" && (
            <Button onClick={handleAddClick} block>Add</Button>
            )}
            </Col> 
            <Col span={2}>
            <Button onClick={handleDelClick} block>Delete</Button>
            </Col> 
           
            <Col span={2}>          
            {managerCheck === "Manager" && (
              <Button id="unlock" onClick={handleUnlockClick} block>Unlock</Button>
            )}
            </Col>
            
          </Row>
          <Row gutter={[40, 0]}>
          <Col span={24}>
          <Table
              columns={columns}
              dataSource={data}
              onRow={(record) => ({
                onClick: () => handleRowClick(record),
              })}
            />

          </Col>
          </Row>
      </div>
    );
}

export default List;