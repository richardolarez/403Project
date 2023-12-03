import React, { useEffect, useState } from 'react';
import {Form, Input, Table, Row, Col, Button, Typography } from 'antd';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

const { Title } = Typography;

const Prescriptions = () => {
  const history = useNavigate();
  const [allData, setAllData] = useState([]);
  const [selectedUser, setSelectedUser] = useState(false);
  const [userData, setUserData] = useState<PresData | null>(null);
  const [roleCheck, setLoggedInUserRole] = useState<string | null>(null);


  interface PresData {
    key: any;
    Drug: any;
    Doses: any;
    CustomerID: any;
    IsFilled: any;
  }

  const [form]=Form.useForm();

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
      IsFilled: pres.IsFilled,
    });
    return data;
  });

  const handleAddClick = () => {
    history('/addPrescription');
  };


  const onFinish = (values: any) => {
    // Call fetchPrescriptions function with form values
    fetch('http://localhost:8080/prescriptions', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(values),
    })
      .then((response) => response.json())
      .then((data) => {
        setAllData(data);
      })
      .catch((error) => {
        console.error('Error fetching prescriptions:', error);
      });
  };

  return (
    <div>
      <Form onFinish={onFinish}>
        <Form.Item name="key" label="Key">
          <Input />
        </Form.Item>
        <Form.Item name="Drug" label="Drug">
          <Input />
        </Form.Item>
        <Form.Item name="Doses" label="Doses">
          <Input />
        </Form.Item>
        <Form.Item name="CustomerID" label="Customer ID">
          <Input />
        </Form.Item>
        <Form.Item name="IsFilled" label="Is Filled">
          <Input />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};
 

export default Prescriptions;
