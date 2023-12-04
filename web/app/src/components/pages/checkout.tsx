import React, { useState, useEffect } from 'react';
import { Form, Button, Input, List } from 'antd';
import { UserOutlined, FieldNumberOutlined, NumberOutlined } from '@ant-design/icons';
import { v4 as uuid } from 'uuid';
import {useNavigate} from 'react-router-dom';
import '../../App.css';
import PaymentPage from './payment';

interface cartItem {
  id: string;
  customerid: number;
  itemid: number;
  paymentMethod: string;
}

const Checkout = () => {
  const history = useNavigate();
  const [cartItems, setCartItems] = useState<cartItem[]>([]);
  const [form] = Form.useForm();
  const [selectedItem, setSelectedItem] = useState<cartItem | null>(null);
  const [paymentMethod, setPaymentMethod] = useState<string>('');

  const addItemToCart = (values: { id: string; customerid: number; itemid: number; paymentMethod: string }) => {
    const itemId = uuid();
    setCartItems([...cartItems, { id: itemId, customerid: values.customerid, itemid: values.itemid, paymentMethod: values.paymentMethod }]);
  };

  const removeItemFromCart = (cartId: string) => {
    if (selectedItem) {
      const updatedCart = cartItems.filter((item) => item !== selectedItem);
      setCartItems(updatedCart);
      setSelectedItem(null);
    }
  };

  const onItemSelected = (item: cartItem) => {
    setSelectedItem(item);
  };

  const onFinish = (values: { id: string; customerid: number; itemid: number; paymentMethod: string }) => {
    console.log('onFinish call');
    fetch('http://localhost:8080/checkout', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ customerid: values.customerid, itemid: values.itemid, paymentMethod: values.paymentMethod, cartItems: cartItems }),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        return response.json();
      })
      .then((cartItem) => {
        console.log('API response');
        setCartItems([...cartItems, { id: values.id, customerid: values.customerid, itemid: values.itemid, paymentMethod: values.paymentMethod }]);
      })
      .catch((error) => {
        console.error('Failed to u', error);
      });
  };

  const handlePayment = (method: string) => {
    setPaymentMethod(method);
    history('/payment')
  };

  useEffect(() => {
    if (paymentMethod) {
      window.location.href = `/payment?method=${paymentMethod}`; // Navigate to payment page when paymentMethod is set
    }
  }, [paymentMethod]);

  return (
    <div>
      <Form name="checkout" initialValues={{ remember: true }} style={{ width: '40%' }} onFinish={onFinish} form={form}>
        <Form.Item name="customerid" label="Customer ID" rules={[{ required: true, message: 'Please enter customer ID' }]}>
          <Input prefix={<UserOutlined />} placeholder="Customer ID" />
        </Form.Item>

        <Form.Item name="itemid" label="Item ID" rules={[{ required: true, message: 'Please enter item ID' }]}>
          <Input prefix={<FieldNumberOutlined />} placeholder="Item ID" />
        </Form.Item>

        <Form.Item name="quantity" label="Qty" rules={[{ required: true, message: 'Please provide a quantity' }]} style={{ width: '100%' }}>
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
          <Button onClick={() => selectedItem && removeItemFromCart(selectedItem.id)} disabled={!selectedItem}>
            Remove Selected Item
          </Button>
        </Form.Item>

        <Form.Item name="paymentMethod" label="Payment Method">
          <Input placeholder="Payment Method" />
        </Form.Item>

        <Form.Item name="signature" label="Signature" rules={[{ required: true, message: 'Please enter your signature' }]}>
          <Input placeholder="Signature" />
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
        renderItem={(item) => (
          <List.Item className={selectedItem && selectedItem.id === item.id ? 'selected-item' : ''} onClick={() => onItemSelected(item)}>
            {`Customer ID: ${item.customerid}, Item ID: ${item.itemid}`}
          </List.Item>
        )}
      />
    </div>
  );
};

export default Checkout;