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

  useEffect(() => {
    axios.get(`http://localhost:8080/prescriptions`).then(res => {
      setAllData(res.data);
  });
  }, []);

  interface PresData {
    ID: number;
    Drug: string;
    Doses: number;
    CustomerID: number;
    IsFilled: any;
    Strength: string;
    Price: number;
    Doctor: string;
    PharmacistID: number;
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

  const data: { Drug: any; Doses: any; CustomerID: any; IsFilled: any; Strength: any; Price: any; Doctor: any; PharmacistID: any;}[] = [];
allData.map((pres: any) => {
  data.push({
    Drug: pres.Drug,
    Doses: pres.Doses,
    CustomerID:  pres.CustomerID,
    IsFilled: pres.IsFilled, 
    Strength: pres.Strength,
    Price: pres.Price,
    Doctor: pres.Doctor,
    PharmacistID: pres.PharmacistID,
  });
  return data;
});



  const handleAddClick = () => {
    history('/addPrescription');
  };


  const onFinish = (values: any) => {
    // Call fetchPrescriptions function with form values
    const payload = {
      id: parseInt(values.id),
      drug: values.drug,
      doses: parseInt(values.doses),
      strength: values.strength,
      price: parseFloat(values.price),
      doctor: values.doctor,
      customerid: parseInt(values.customerID),
      pharmacistid: parseInt(values.pharmacistID),
      isfilled: "Yes",
    };

    axios.post('http://localhost:8080/prescription', payload, {
      headers: {
        'Content-Type': 'application/json'
      }
    })
  };

  return (
    <div>
      <Form onFinish={onFinish}>
        <Form.Item name="id" label="Prescription ID">
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
        <Form.Item name="pharmacistID" label="Pharmacist ID">
          <Input />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit">
            Fill Prescription
          </Button>
        </Form.Item>
      </Form>

      <Table columns={columns} dataSource={data} />
    </div>
  );
};
 

export default Prescriptions;
