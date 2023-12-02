import React, { useState } from 'react';
import { Form, Radio, Button } from 'antd';

const PaymentPage: React.FC = () => {
    const [paymentMethod, setPaymentMethod] = useState('');

    const handlePaymentSubmit = (values: any) => {
        // Handle payment submission logic here
        console.log('Payment submitted:', values.paymentMethod);
    };

    return (
        <div>
            <h1>Payment Page</h1>
            <Form onFinish={handlePaymentSubmit}>
                <Form.Item name="paymentMethod">
                    <Radio.Group>
                        <Radio value="cash">Cash</Radio>
                        <Radio value="credit">Credit</Radio>
                        <Radio value="debit">Debit</Radio>
                    </Radio.Group>
                </Form.Item>
                <br />
                <Form.Item>
                    <Button type="primary" htmlType="submit">
                        Submit Payment
                    </Button>
                </Form.Item>
            </Form>
        </div>
    );
};

export default PaymentPage;
