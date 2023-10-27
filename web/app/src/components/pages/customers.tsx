//src/components/pages/customers.tsx
import React, {useEffect, useState} from 'react';
import {Table, Row, Col, Button, Typography} from 'antd';
import {useNavigate} from 'react-router';
import axios from 'axios';

const {Title} = Typography;


const Customers = () => {
  const history = useNavigate();
  const [allData, setAllData] = useState([]);
  const [selectedUser, setSelectedUser] = useState(false);
  const [userData, setUserData] = useState<CustomerData | null>(null);
  const [managerCheck, setLoggedInUserRole] = useState<string | null>(null);



  useEffect(() => {
      axios.get(`http://localhost:8080/customers`).then(res => {
        setAllData(res.data);
      });
    }, []);

  interface CustomerData {
    key: any;
    FirstName: any;
    LastName: any;
    Email: any;
    PhoneNumber: any;
    Address: any;
  }

  const columns = [
    {
      title: 'First Name',
      dataIndex: 'FirstName',
    },
    {
      title: 'Last Name',
      dataIndex: 'LastName'
    },
    {
      title: 'Email',
      dataIndex: 'Email'
    },
    {
      title: 'Phone Number',
      dataIndex: 'PhoneNumber'
    },
    {
      title: 'address',
      dataIndex: 'Address'
    }
  ];

  const data: {
    key: any;
    FirstName: any;
    LastName: any;
    Email: any;
    PhoneNumber: any;
    Address: any;
  }[] = [];
    allData.map((user: any) => {
        data.push({
        key: user.ID,
        FirstName: user.FirstName,
        LastName:  user.LastName,
        Email:     user.Email,
        PhoneNumber: user.PhoneNumber,
        Address:    user.Address
      })
      return data;
    });
  console.log(data)

  const handleAddClick = () => {
    history('/custForm')
    }
  
  const handleDelClick = () => {
    if (!userData) {
      console.error('No user selected for deletion');
      return;
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
  
  const handleModClick = () => {
    history('/form')
    }

  const handleRowClick = (record : CustomerData) => {
    setSelectedUser(record.FirstName) // using horrible JS true value interpretation
    setUserData(record);
  }

    
  return (
      <div>
          <Row gutter={[40, 0]}>
            <Col span={10}>
              <Title level={2}>
              Customers List
              </Title>
              </Col>
            <Col span={15}>
              {selectedUser && <Title level={3}>Selected User: {selectedUser}</Title>}
            </Col>
            <Col span={2}>
            <Button onClick={handleAddClick} block>Add</Button>
            </Col> 
            <Col span={2}>
            <Button onClick={handleDelClick} block>Delete</Button>
            </Col> 
            <Col span={2}>
            <Button onClick={handleModClick} block>Modify</Button>
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

export default Customers;