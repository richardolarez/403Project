import React, { useState } from 'react';
import { Typography, message } from 'antd';
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
            <label>
                Name:
                <input
                    type="text"
                    name="name"
                    value={formData.name}
                    onChange={handleChange}
                />
            </label>
            <label>
                Quantity:
                <input
                    type="number"
                    name="quantity"
                    value={formData.quantity}
                    onChange={handleChange}
                />
            </label>
            <label>
                Price:
                <input
                    type="number"
                    name="price"
                    value={formData.price}
                    onChange={handleChange}
                />
            </label>
            <label>
                Name:
                <input
                    type="text"
                    name="name"
                    value={formData.name}
                    onChange={handleChange}
                />
            </label>
            <label>
                Quantity:
                <input
                    type="number"
                    name="quantity"
                    value={formData.quantity}
                    onChange={handleChange}
                />
            </label>
            <label>
                Price:
                <input
                    type="number"
                    name="price"
                    value={formData.price}
                    onChange={handleChange}
                />
            </label>
            <button type="submit">Submit</button>
        </form>
    );
}

export default Inventory;
