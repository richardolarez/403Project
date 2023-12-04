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
    Strength: any;
    Price: any;
    Doctor: any;
    PharmacistID: any;
  }

  const [form] = Form.useForm();

  const columns = [
    {
      title: 'Drug',
      dataIndex: 'Drug',
    },
    {
      title: 'Doses',
      dataIndex: 'Doses',
    },
    {
      title: 'Customer ID',
      dataIndex: 'CustomerID',
    },
    {
      title: 'Filled?',
      dataIndex: 'IsFilled',
    },
    {
      title: 'Strength',
      dataIndex: 'Strength',
    },
    {
      title: 'Price',
      dataIndex: 'Price',
    },
    {
      title: 'Doctor',
      dataIndex: 'Doctor',
    },
    {
      title: 'Pharmacist ID',
      dataIndex: 'PharmacistID',
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
    axios
    .post('http://localhost:8080/prescription', values)
    .then((res) => {
      console.log(res);
      })
      .catch((error) => {
        console.log(error);
      });
  };

  return (
    <div>
      <Form onFinish={onFinish}>
        <Form.Item name="id" label="ID">
          <Input />
        </Form.Item>
        <Form.Item name="drug" label="Drug">
          <Input />
        </Form.Item>
        <Form.Item name="doses" label="Doses">
          <Input />
        </Form.Item>
        <Form.Item name="strength" label="Strength">
          <Input />
        </Form.Item>
        <Form.Item name="price" label="Price">
          <Input />
        </Form.Item>
        <Form.Item name="doctor" label="Doctor">
          <Input />
        </Form.Item>
        <Form.Item name="customerID" label="Customer ID">
          <Input />
        </Form.Item>
        <Form.Item name="isFilled" label="Is Filled">
          <Input />
        </Form.Item>
        <Form.Item name="pharmacistID" label="Pharmacist ID">
          <Input />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
        </Form.Item>
      </Form>

      <Table columns={columns} dataSource={data} />
    </div>
  );
};
 

export default Prescriptions;
