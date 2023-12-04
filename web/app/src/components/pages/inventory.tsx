import React, { useState } from 'react';
import { Typography, message, Input, Button, Row, Col, Form, Checkbox, InputNumber } from 'antd';
import axios from 'axios';
import { useNavigate } from 'react-router';

const { Title } = Typography;

const layout = {
    labelCol: { span: 8 },
    wrapperCol: { span: 16 },
  };

interface InventoryFormProps {
    onSubmit: (formData: FormData) => void;
    initialFormData?: FormData;
}

const Inventory = () => {

    const [loading, setLoading] = useState(false);
    const history = useNavigate();


    const handleSubmit = (values: any) => {
        setLoading(true);
        axios
            .post('http://localhost:8080/addNewInventoryItem', values)
            .then((res) => {
                setLoading(false);
                message.success('Updated Item!');
                history('/inventoryList');
            })
            .catch((error) => {
                setLoading(false);
                message.error(error.message);
            });
    };

    return (
        <div>
        <Row gutter={[40, 0]}>
          <Col span={23}>
            <Title style={{textAlign: 'center'}} level={2}>
            Input New item
            </Title>
            </Col>
        </Row>
        <Row gutter={[40, 0]}>
        <Col span={18}>
        <Form
        {...layout}
          name="Inventory form"
          onFinish={handleSubmit}
        >
              <Form.Item label="Name" name="Name">
                <Input type="text" />
              </Form.Item>
              <Form.Item label="Description" name="Description">
                <Input type="text" />
              </Form.Item>
              <Form.Item label="Price" name="Price">
              <InputNumber
                            min={0}
                            step={1}
                            style={{ width: '100%' }}
                            parser={(value: string | undefined) => parseInt(value || '0') || 0}
                            formatter={(value: number | undefined) => `${value}`}
                        />
              </Form.Item>
              <Form.Item label="Quantity" name="Quantity">
              <InputNumber
                            min={0}
                            step={1.0}
                            style={{ width: '100%' }}
                            parser={(value: string | undefined) => parseFloat(value || '0') || 0}
                            formatter={(value: number | undefined) => `${value}`}
                        />
              </Form.Item>
              
     
          <div style={{textAlign: "right"}}>
            {
              <Button
                type="primary"
                loading={loading}
                htmlType="submit"
              >
                Save
              </Button>
          }
            <Button type="default" htmlType="button" onClick={() => history('/inventoryList')}>
              Back
            </Button>
              </div>
        </Form>
        </Col>
        </Row>
        </div>
      );
}

export default Inventory;
