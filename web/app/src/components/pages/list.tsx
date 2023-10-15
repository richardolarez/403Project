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

  useEffect(() => {
      axios.get(`http://localhost:8080/employees`).then(res => {
        setAllData(res.data);
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
    history('/form')
    }
  
  const handleModClick = () => {
    history('/form')
    }
  const handleRowClick = (record : UserData) => {
    setSelectedUser(record.Username) // using horrible JS true value interpretation
    
  }
  
  return (
      <div>
          <Row gutter={[40, 0]}>
            <Col span={8}>
              <Title level={2}>
              User List
              </Title>
              </Col>
            <Col span={10}>
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

export default List;