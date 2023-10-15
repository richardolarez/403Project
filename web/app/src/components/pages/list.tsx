//src/components/pages/list.tsx
import React, {useEffect, useState} from 'react';
import {Table, Row, Col, Button, Typography} from 'antd';
import {useNavigate} from 'react-router';
import axios from 'axios';

const {Title} = Typography;

const List = () => {
  const history = useNavigate();
  const [allData, setAllData] = useState([]);

  useEffect(() => {
      axios.get(`http://localhost:8080/employees`).then(res => {
        setAllData(res.data);
      });
    }, []);

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

  const handleClick = () => {
    history('/form')
    }
  
  return (
      <div>
          <Row gutter={[40, 0]}>
            <Col span={18}>
              <Title level={2}>
              User List
              </Title>
              </Col>
            <Col span={6}>
            <Button onClick={handleClick} block>Add User</Button>
            </Col>
          </Row>
          <Row gutter={[40, 0]}>
          <Col span={24}>
          <Table columns={columns} dataSource={data} />
          </Col>
          </Row>
      </div>
    );
}

export default List;