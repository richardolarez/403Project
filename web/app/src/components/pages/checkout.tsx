import React, { useState, useEffect } from 'react';
import { Card, Form, Button, Input, List, message } from 'antd';
import { UserOutlined, FieldNumberOutlined, NumberOutlined } from '@ant-design/icons';
import { v4 as uuid } from 'uuid';
import { useNavigate } from 'react-router-dom';
import '../../App.css';
import PaymentPage from './payment';

interface cartItem {
  id: string;
  customerid: string;
  itemid: string;
  paymentMethod: string;
}

interface InventoryItem {
  ID: number;
  Name: string;
  Description: string;
  Quantity: number;
  Price: number;
  IsPrescription: boolean;
}

interface SalesReceipt {
  TransactionID: number;
  Items: InventoryItem[];
  CustomerID: number;
  TotalAmount: number;
  PaymentMethod: string;
  // Add more fields as needed
}

  const Checkout = () => {
    const [form] = Form.useForm();
    const [selectedItem, setSelectedItem] = useState<cartItem | null>(null);
    const [paymentMethod, setPaymentMethod] = useState<string>('');
    const [receipt, setReceipt] = useState<SalesReceipt | null>(null);
    const [cartItems, setCartItems] = useState<cartItem[]>([]);

    const addItemToCart = (values: { id: string; customerid: string; itemid: string; paymentMethod: string }) => {
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

    const onFinish = (values: { id: string; customerid: string; itemid: string; paymentMethod: string }) => {
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
          return response.json(); // Modify this line to return response as JSON
        })
        .then((receipt: SalesReceipt) => {
          console.log('API response');
          setCartItems([...cartItems, { id: values.id, customerid: values.customerid, itemid: values.itemid, paymentMethod: values.paymentMethod }]);
          setReceipt(receipt); // Set the receipt state
        })
        .catch((error) => {
          console.error('Failed to u', error);
        });
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
        {receipt && (
          <Card title="Receipt">
            <p>Transaction ID: {receipt.TransactionID}</p>
            <p>Customer ID: {receipt.CustomerID}</p>
            <p>Total Amount: {receipt.TotalAmount}</p>
            <p>Payment Method: {receipt.PaymentMethod}</p>
            <p>Items:</p>
          </Card>
        )}
      </div>
    );
  };


  export default Checkout;