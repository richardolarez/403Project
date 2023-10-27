//src/components/pages/form.tsx
import React, {useState, useEffect} from 'react';
import {Row, Col, Typography, Input, Form, Button, 
Radio, Switch, Slider, Select, message, Checkbox, Tooltip} from 'antd';
import axios from 'axios';
import {useNavigate} from 'react-router';
const {Title} = Typography;
const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};
const FormApp = () => {
  const [loading, setLoading] = useState(false);
  const [isCustomer, setIsCustomer] = useState(false);
  const [userRole, setUserRole] = useState('');

  const history = useNavigate();
  
const handleSubmit = (values: any) => {
    setLoading(true);

    if (isCustomer) {
      values.role = 'Customer';
    }

    axios.post(`http://localhost:8080/addEmployee`, 
      values
    )
    .then(res => {
      setLoading(false);
      message.success('User Added Successfully!');
      history('/list');
    })
    .catch(error => {
      setLoading(false);
      message.error(error);
    })
  }

  const roleRules = isCustomer
    ? 
    [] : [
      {
        required: true,
        message: 'Please select your Role',
      },
    ];

    useEffect(() => {
      // Load user role from session storage
      const storedUserRole = sessionStorage.getItem('UserRole');
      if (storedUserRole) {
        setUserRole(storedUserRole);
      }
    }, []);
    
return (
    <div>
        <Row gutter={[40, 0]}>
          <Col span={23}>
            <Title style={{textAlign: 'center'}} level={2}>
            Please Fill the User Form
            </Title>
            </Col>
        </Row>
        <Row gutter={[40, 0]}>
        <Col span={18}>
          <Form {...layout} onFinish={handleSubmit}>
            <Form.Item name="username" label="UserName"
            rules={[
              {
                required: true,
                message: 'Please input your name',
              }
            ]}
            >
              <Input placeholder="Please Enter your username" />
            </Form.Item>
            <Form.Item name="password" label="password" 
            rules={[
              {
                required: true,
                message: 'Please input your password'
              }
            ]}
            >
              <Input placeholder="Please Enter your desired password" />
            </Form.Item>
            <Form.Item name="firstname" label="First Name" 
            rules={[
              {
                required: true,
                message: 'Please input your First name',
              }
            ]}
            >
              <Input placeholder="Please Enter your First Name" />
            </Form.Item>
            <Form.Item name="lastname" label="Last Name" 
            rules={[
              {
                required: true,
                message: 'Please input your Last name',
              }
            ]}
            >
              <Input placeholder="Please Enter your Last Name" />
            </Form.Item>
            <Form.Item name="role" label="Role" 
            rules={roleRules}
            >
              <Select  placeholder="Please select your role" disabled={isCustomer}>
                <Select.Option value="Cashier">Cashier</Select.Option>
                <Select.Option value="Manager">Manager</Select.Option>
                <Select.Option value="Pharmacist">Pharmacist</Select.Option>
                <Select.Option value="BadGuy">Martin Shkreli</Select.Option>
              </Select>
            </Form.Item>

            <Form.Item name="isCustomer" label="Is Customer">
              <Checkbox
                checked={isCustomer}
                onChange={(e) => setIsCustomer(e.target.checked)}
              >
                Is Customer
              </Checkbox>
              
            </Form.Item>

            <div style={{textAlign: "right"}}>
            {userRole !== 'Manager' && (
            <Tooltip
              title="You is NOT a manager!"
              placement="bottom"
            >
              <Button
                type="primary"
                loading={loading}
                htmlType="submit"
                disabled={userRole !== 'Manager'}
              >
                Save
              </Button>
            </Tooltip>
          )}
            <Button type="default" htmlType="button" onClick={() => history('/list')}>
              Back
            </Button>
              </div>
          </Form>
          </Col>
        </Row>
    </div>
  );
}
export default FormApp;