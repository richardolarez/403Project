import React, { useState } from 'react';
import { Typography, message, Input, Button, Row, Col } from 'antd';
import axios from 'axios';
import { useNavigate } from 'react-router';

const { Title } = Typography;

interface InventoryFormProps {
    onSubmit: (formData: FormData) => void;
    initialFormData?: FormData;
}

const Inventory = () => {
    const [formData, setFormData] = useState({
        name: '',
        quantity: 0,
        price: 0,
    });
    const [loading, setLoading] = useState(false);
    const history = useNavigate();

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        setFormData({ ...formData, [name]: value });
    };

    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setLoading(true);
        axios
            .post('http://localhost:8080/updateInventoryItem', formData)
            .then((res) => {
                setLoading(false);
                message.success('Updated Item!');
                history('/list');
            })
            .catch((error) => {
                setLoading(false);
                message.error(error.message);
            });
    };

    return (
        <form onSubmit={handleSubmit}>
            <Row gutter={16}>
                <Col span={8}>
                    <label>
                        Name:
                        <Input
                            type="text"
                            name="name"
                            value={formData.name}
                            onChange={handleChange}
                        />
                    </label>
                </Col>
                <Col span={8}>
                    <label>
                        Quantity:
                        <Input
                            type="number"
                            name="quantity"
                            value={formData.quantity}
                            onChange={handleChange}
                        />
                    </label>
                </Col>
                <Col span={8}>
                    <label>
                        Price:
                        <Input
                            type="number"
                            name="price"
                            value={formData.price}
                            onChange={handleChange}
                        />
                    </label>
                </Col>
            </Row>
            <Button type="primary" htmlType="submit" loading={loading}>
                Submit
            </Button>
        </form>
    );
}

export default Inventory;
