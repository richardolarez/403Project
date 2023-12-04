import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Table, Typography, Row, Col, Button, Modal, Form, Input, Checkbox,
InputNumber } from 'antd';
import { useNavigate } from 'react-router';

const { Title } = Typography;

interface InventoryItem {
    ID: number;
    Name: string;
    Description: string;
    Quantity: number;
    Price: number;
    IsPrescription: boolean;
}

interface Medicine {
    ID: number;
    Drug: string;
    Doses: number;
    Price: number;
    ExpirationDate: Date;
}

const InventoryList: React.FC = () => {
    const history = useNavigate();
    const [inventoryItems, setInventoryItems] = useState<InventoryItem[]>([]);
    const [selectedItem, setSelectedItem] = useState(false);
    const [item, setItem] = useState<InventoryItem | null>(null);
    const [medicines, setMedicines] = useState<Medicine[]>([]);
    const [updateModalVisible, setUpdateModalVisible] = useState(false);
    const [form] = Form.useForm();

    useEffect(() => {
        const fetchInventoryItems = async () => {
            try {
                const response = await axios.get('http://localhost:8080/inventory');
                if (response.status === 200) {
                    setInventoryItems(response.data);
                } else {
                    console.error('Failed to fetch inventory items');
                }
                const medicineResponse = await axios.get('http://localhost:8080/medicines');
                if (medicineResponse.status === 200) {
                    setMedicines(medicineResponse.data);
                } else {
                    console.error('Failed to fetch medicine data');
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
       
    ];

    const medicineColumns = [
        {
            title: 'ID',
            dataIndex: 'ID',
        },
        {
            title: 'Drug',
            dataIndex: 'Drug',
        },
        {
            title: 'Doses',
            dataIndex: 'Doses',
        },
        {
            title: 'Strength',
            dataIndex: 'Strength',
        },
        {
            title: 'Price',
            dataIndex: 'Price',
        },
        {
            title: 'ExpirationDate',
            dataIndex: 'ExpirationDate',
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


    const handleAddClick = () => {
        history('/inventory')
    }

    const handleModalOk = (values: any) => {
        if (!item) {
            console.error('No  selected item for update');
            return;
          } 

          const updateData = {
            ID: item.ID,
            ...values
          }
          console.log(updateData)
        axios.post(`http://localhost:8080/updateInventoryItem`, updateData).then(res => {
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
                    <Button block onClick={handleAddClick}>Add</Button>
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
            <Row gutter={[40, 0]}>
                <Col span={24}>
                <Title level={3}>Medications</Title>
                    <Table
                        columns={medicineColumns}
                        dataSource={medicines}
                    />
                </Col>
            </Row>
            <Modal
                title="Update Item"
                visible={updateModalVisible}
                onCancel={handleModalCancel}
                onOk={() => form.submit()}
            >
                <Form form ={form} layout="vertical" onFinish={handleModalOk}>
                    <Form.Item label="Name" name="Name">
                        <Input />
                    </Form.Item>
                    <Form.Item label="Description" name="Description">
                        <Input />
                    </Form.Item>
                    <Form.Item label="Quantity" name="Quantity">
                        <InputNumber
                            min={0}
                            step={1}
                            style={{ width: '100%' }}
                            parser={(value: string | undefined) => parseInt(value || '0') || 0}
                            formatter={(value: number | undefined) => `${value}`}
                        />
                    </Form.Item>
                    <Form.Item label="Price" name="Price">
                        <InputNumber
                            min={0}
                            step={1.0}
                            style={{ width: '100%' }}
                            parser={(value: string | undefined) => parseFloat(value || '0') || 0}
                            formatter={(value: number | undefined) => `${value}`}
                        />
                    </Form.Item>
                </Form>
            </Modal>

        </div>
    );
};

export default InventoryList;
