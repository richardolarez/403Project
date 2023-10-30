import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Table, Typography, Row, Col, Button, Modal, Form, Input, Checkbox,
InputNumber } from 'antd';

const { Title } = Typography;

interface InventoryItem {
    Id: number;
    Name: string;
    Description: string;
    Quantity: number;
    Price: number;
    IsPrescription: boolean;
}

const InventoryList: React.FC = () => {
    const [inventoryItems, setInventoryItems] = useState<InventoryItem[]>([]);
    const [selectedItem, setSelectedItem] = useState(false);
    const [item, setItem] = useState<InventoryItem | null>(null);
    const [updateModalVisible, setUpdateModalVisible] = useState(false);

    useEffect(() => {
        const fetchInventoryItems = async () => {
            try {
                const response = await axios.get('http://localhost:8080/inventory');
                if (response.status === 200) {
                    setInventoryItems(response.data);
                } else {
                    console.error('Failed to fetch inventory items');
                }
            } catch (error) {
                console.error('Failed to fetch inventory items:', error);
            }
        };

        fetchInventoryItems();
    }, []);

    const columns = [
        {
            title: 'Name',
            dataIndex: 'Name',
        },
        {
            title: 'Description',
            dataIndex: 'Description',
        },
        {
            title: 'Quantity',
            dataIndex: 'Quantity',
        },
        {
            title: 'Price',
            dataIndex: 'Price',
        },
        {
            title: 'Prescription?',
            dataIndex: 'IsPrescription',
            render: (IsPrescription: boolean) => (
                IsPrescription ? '✓' : 'X'
            ),
        },
    ];

    const handleRowClick = (record: InventoryItem) => {
        setSelectedItem(true);
        setItem(record);
        console.log(record)
    };

    const handleUpdateClick = () => {
        setUpdateModalVisible(true);
        console.log(item)
    };

    const handleModalOk = () => {
        if (!item) {
            console.error('No  selected item for update');
            return;
          } 

          const updateData = {
            ID: item.Id,
            Name: item.Name,
            Description: item.Description,
            Price: item.Price,
            Quantity: item.Quantity,
            IsPrescription: item.IsPrescription
          }
          console.log(updateData)
        axios.post(`http://localhost:8080/updateInventoryItem`, {
            data: updateData,
        }).then(res => {
            axios.get('http://localhost:8080/inventory').then((res) => {
              setInventoryItems(res.data);
            });
            setSelectedItem(false)
          })
          .catch((error) => {
            console.error('Error deleting employee:', error);
          });
        setUpdateModalVisible(false);
    };

    const handleModalCancel = () => {
        setUpdateModalVisible(false);
    };

    return (
        <div>
            <Row gutter={[40, 0]}>
                <Col span={10}>
                    <Title level={2}>Inventory List</Title>
                </Col>
                <Col span={18}>
                    {selectedItem && <Title level={3}>Selected Item: {item?.Name}</Title>}
                </Col>
                <Col span={2}>
                    <Button block>Add</Button>
                </Col>
                <Col span={2}>
                    <Button block onClick={handleUpdateClick}>Update</Button>
                </Col>
            </Row>
            <Row gutter={[40, 0]}>
                <Col span={24}>
                    <Table
                        columns={columns}
                        dataSource={inventoryItems}
                        onRow={(record) => ({
                            onClick: () => handleRowClick(record),
                        })}
                    />
                </Col>
            </Row>
            <Modal
                title="Update Item"
                visible={updateModalVisible}
                onOk={handleModalOk}
                onCancel={handleModalCancel}
            >
                <Form layout="vertical">
                    <Form.Item label="Name">
                        <Input />
                    </Form.Item>
                    <Form.Item label="Description">
                        <Input />
                    </Form.Item>
                    <Form.Item label="Quantity">
                        <InputNumber
                            min={0}
                            step={1}
                            style={{ width: '100%' }}
                            parser={(value: string | undefined) => parseInt(value || '0') || 0}
                            formatter={(value: number | undefined) => `${value}`}
                        />
                    </Form.Item>
                    <Form.Item label="Price">
                        <InputNumber
                            min={0}
                            step={1.0}
                            style={{ width: '100%' }}
                            parser={(value: string | undefined) => parseFloat(value || '0') || 0}
                            formatter={(value: number | undefined) => `${value}`}
                        />
                    </Form.Item>
                    <Form.Item label="Prescription">
                        <Checkbox />
                    </Form.Item>
                </Form>
            </Modal>

        </div>
    );
};

export default InventoryList;
