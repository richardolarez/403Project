import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

interface InventoryFormProps {
    onSubmit: (formData: FormData) => void;
    initialFormData?: FormData;
}

interface FormData {
    name: string;
    quantity: number;
    price: number;
}

const Inventory: React.FC<InventoryFormProps> = ({ onSubmit, initialFormData }) => {
    const [formData, setFormData] = useState<FormData>(
        initialFormData || { name: '', quantity: 0, price: 0 }
    );
    const history = useNavigate();

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        setFormData({ ...formData, [name]: value });
    };

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const response = await fetch(initialFormData ? '/updateInventoryItem' : '/addNewInventoryItem', {
            method: initialFormData ? 'PUT' : 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        });
        if (response.ok) {
            onSubmit(formData);
            navigate('/inventory'); // navigate to inventory page after successful submission
        } else {
            console.error('Failed to update inventory item');
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <label>
                Name:
                <input type="text" name="name" value={formData.name} onChange={handleChange} />
            </label>
            <label>
                Quantity:
                <input type="number" name="quantity" value={formData.quantity} onChange={handleChange} />
            </label>
            <label>
                Price:
                <input type="number" name="price" value={formData.price} onChange={handleChange} />
            </label>
            <button type="submit">Submit</button>
        </form>
    );
};

export default Inventory;
