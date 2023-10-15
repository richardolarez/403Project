import React from 'react';
import { Form, Button, Input } from 'antd';
import { UserOutlined, FieldNumberOutlined,  NumberOutlined } from '@ant-design/icons';


const Checkout = () => {

  const onFinish = (values: { customerid: number, itemid: number, quantity: number }) => {
    // Call the /login endpoint to authenticate the employee
    // fetch('http://localhost:8080/checkout', {
    //   method: 'POST',
    //   headers: {
    //     'Content-Type': 'application/json'
    //   },
    //   body: JSON.stringify({ username: values.username, password: values.password })
    // })
    // .then(response => {
    //   if (!response.ok) {
    //     throw new Error(`HTTP error! status: ${response.status}`);
    //   }
    //   return response.json();
    // })
    // .then(employee => {
    //   console.log('Authenticated employee:', employee);
    //   sessionStorage.setItem('authenticated', 'true');
    //   sessionStorage.setItem('UserFName', employee.FirstName);
    //   window.location.reload();
    // })
    // .catch(error => {
    //   console.error('Failed to u', error);
    //   // TODO: Handle authentication error
    // });
  };

      return (
        <div>
          <Form
          name="checkout"
          initialValues={{ remember: true }}
          style={{ width: '40%'}}
          >
          <Form.Item
            name="customerid" label="Customer ID"
            rules={[{ required: true, message: 'Please enter customer ID' }]}
          >
            <Input prefix={<UserOutlined />} placeholder="Customer ID" />
          </Form.Item>

          <Form.Item
            name="itemid" label="Item ID"
            rules={[{ required: true, message: 'Please enter item ID' }]}
          >
            <Input prefix={<FieldNumberOutlined />} placeholder="Item ID" />
          </Form.Item>

          <Form.Item
            name="qty" label="Qty"
            rules={[{ required: true, message: 'Please provide a quantity' }]}
            style={{ width: '100%'}}
          >
            <div style={{ display: 'flex', alignItems: 'center' }}>
              <Input prefix={<NumberOutlined />} placeholder="Qty" />
              <Button style={{ marginLeft: 10 }}>+</Button>
            </div>
          </Form.Item>

          <Form.Item>
            <Button type="default" htmlType="submit" style={{ width: '100%' }}>
              Add Item
            </Button>
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" style={{ width: '100%' }}>
              Checkout
            </Button>
          </Form.Item>

          </Form>
        </div>

      );
    };

export default Checkout;