//src/components/pages/list.tsx
import React, {useEffect, useState} from 'react';
import {Table, Row, Col, Button, Typography} from 'antd';
import {useNavigate} from 'react-router';
import axios from 'axios';

const {Title} = Typography;


const Prescriptions = () => {
  const history = useNavigate();
  const [allData, setAllData] = useState([]);
  const [selectedUser, setSelectedUser] = useState(false);
  const [userData, setUserData] = useState<PresData | null>(null);
  //Retrieve shadow realm from localstorage.
  const [lockedAccounts, setLockedAccounts] = useState<string[]>(JSON.parse(localStorage.getItem('shadowRealm') || '[]'));
  const [managerCheck, setLoggedInUserRole] = useState<string | null>(null);



  useEffect(() => {
      axios.get(`http://localhost:8080/prescriptions`).then(res => {
        setAllData(res.data);
        setLoggedInUserRole(sessionStorage.getItem("UserRole"));
        //console.log("session role: " + sessionStorage.getItem("UserRole"))
      });
    }, []);

  interface PresData {
    key: any;
    Drug: any;
    Doses: any;
    CustomerID: any;
    IsFilled: any;
  }

  const columns = [
    {
      title: 'Drug',
      dataIndex: 'Drug',
    },
    {
      title: 'Doses',
      dataIndex: 'Doses'
    },
    {
      title: 'Customer ID',
      dataIndex: 'CustomerID'
    },
    {
      title: 'Filled?',
      dataIndex: 'IsFilled'
    },
  ];

  const data: { key: any; Drug: any; Doses: any; CustomerID: any; IsFilled: any; }[] = [];
    allData.map((pres: any) => {
        data.push({
        key: pres.ID,
        Drug: pres.Drug,
        Doses: pres.Doses,
        CustomerID:  pres.CustomerID,
        IsFilled:      pres.IsFilled,
      })
      return data;
    });
  console.log(data)

  const handleAddClick = () => {
    history('/form')
    }
  
  const handleDelClick = () => {
    if (!userData) {
      console.error('No prescription selected for deletion');
      return;
    }

    const deleteData = {
      id: userData.key,

    }

    axios.delete(`http://localhost:8080/deletePrescription`, {
      data: deleteData,
    })
    .then(res => {
      axios.get('http://localhost:8080/prescriptions').then((res) => {
        setAllData(res.data);
      });
    })
    .catch((error) => {
      console.error('Error deleting prescription:', error);
    });
    }
  
  const handleModClick = () => {
    history('/form')
    }

  const handleRowClick = (record : PresData) => {
    setSelectedUser(record.key) // using horrible JS true value interpretation
    setUserData(record);
  }

  const handleUnlockClick = () => {
    if (!userData) {
      console.error('No user selected');
      return;
    }
  };
    
  return (
      <div>
          <Row gutter={[40, 0]}>
            <Col span={10}>
              <Title level={2}>
              Prescriptions
              </Title>
              </Col>
            <Col span={15}>
              {selectedUser && <Title level={3}>Selected Prescription: {selectedUser}</Title>}
            </Col>
            <Col span={2}>
            <Button onClick={handleAddClick} block>Add</Button>
            </Col> 
            <Col span={2}>
            <Button onClick={handleModClick} block>Fill</Button>
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



export default Prescriptions;