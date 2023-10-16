import React, { useState } from 'react';
import { Form, Button, Input, List } from 'antd';
import { UserOutlined, FieldNumberOutlined,  NumberOutlined } from '@ant-design/icons';

interface cartItem{
customerid: number;
itemid: number;
quantity: number;
}



const Checkout = () => {

  const [cartItems, setCartItems] = useState<cartItem[]>([]);
  const [form] = Form.useForm();

  const addItemToCart = (values: { customerid: number, itemid: number, quantity: number }) => {
    setCartItems([...cartItems, {
        customerid: values.customerid,
        itemid: values.itemid,
        quantity: values.quantity
    }]);
};

  const onFinish = (values: { customerid: number, itemid: number, quantity: number }) => {
    console.log("onFinish call");
    fetch('http://localhost:8080/checkout', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ customerid: values.customerid, itemid: values.itemid, quantity: values.quantity })
    })
    .then(response => {
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      return response.json();
    })
    .then(cartItem => {
      // Add the item to the cart items state
      console.log("API response");
      setCartItems([...cartItems,{customerid: values.customerid, itemid: values.itemid, quantity: values.quantity}]);
    })
    .catch(error => {
      console.error('Failed to u', error);
      // TODO: Handle authentication error
    });
  };

      return (
        <div>
          <Form
          name="checkout"
          initialValues={{ remember: true }}
          style={{ width: '40%'}}
          onFinish={onFinish}
          form={form}
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
            name="quantity" label="Qty"
            rules={[{ required: true, message: 'Please provide a quantity' }]}
            style={{ width: '100%'}}
          >
            <div style={{ display: 'flex', alignItems: 'center' }}>
              <Input prefix={<NumberOutlined />} placeholder="Qty" />
              <Button style={{ marginLeft: 10 }}>+</Button>
            </div>
          </Form.Item>

          <Form.Item>
          <Button type="default" onClick={() => addItemToCart(form.getFieldsValue())} style={{ width: '100%' }}>
              Add Item
            </Button>
          </Form.Item>

          <Form.Item>
          <Button type="default" htmlType="submit" style={{ width: '100%' }}>
              Checkout
            </Button>
          </Form.Item>

          </Form>
          <List
        header={<div>Cart Items</div>}
        bordered
        dataSource={cartItems}
        renderItem={item => (
          <List.Item>
            {`Customer ID: ${item.customerid}, Item ID: ${item.itemid}, Quantity: ${item.quantity}`}
          </List.Item>
        )}
      />
        </div>

      );
    };

export default Checkout;