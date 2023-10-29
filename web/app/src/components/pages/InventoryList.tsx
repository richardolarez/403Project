import React, { useState, useEffect } from 'react';

interface InventoryItem {
    id: number;
    name: string;
    quantity: number;
    price: number;
}

const InventoryList: React.FC = () => {
    const [inventoryItems, setInventoryItems] = useState<InventoryItem[]>([]);

    useEffect(() => {
        const fetchInventoryItems = async () => {
            const response = await fetch('/getInventoryItems');
            if (response.ok) {
                const data = await response.json();
                setInventoryItems(data);
            } else {
                console.error('Failed to fetch inventory items');
            }
        };
        fetchInventoryItems();
    }, []);

    return (
        <div>
            <h1>Inventory List</h1>
            <table>
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Quantity</th>
                        <th>Price</th>
                    </tr>
                </thead>
                <tbody>
                    {inventoryItems.map((item) => (
                        <tr key={item.id}>
                            <td>{item.name}</td>
                            <td>{item.quantity}</td>
                            <td>{item.price}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default InventoryList;