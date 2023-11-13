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
const CustForm = () => {
  const [loading, setLoading] = useState(false);
  const [isCustomer, setIsCustomer] = useState(false);
  const [userRole, setUserRole] = useState('');

  const history = useNavigate();
  
const handleSubmit = (values: any) => {
    setLoading(true);
    const logout = function() {
      sessionStorage.removeItem('authenticated');
      sessionStorage.removeItem("UserFname");
      sessionStorage.removeItem("UserRole");
  
      window.location.href = '/login';
     }


    axios.post(`http://localhost:8080/addCustomer`, 
      values
    )
    .then(res => {
      setLoading(false);
      message.success('Customer Added Successfully!');
      history('/customers');
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
            Please Fill the Customer Form
            </Title>
            </Col>
        </Row>
        <Row gutter={[40, 0]}>
        <Col span={18}>
          <Form {...layout} onFinish={handleSubmit}>
            <Form.Item name="FirstName" label="First Name"
            rules={[
              {
                required: true,
                message: 'Please input your name',
              }
            ]}
            >
              <Input placeholder="Please Enter First Name" />
            </Form.Item>
            <Form.Item name="LastName" label="Last Name" 
            rules={[
              {
                required: true,
                message: 'Please input your Last name'
              }
            ]}
            >
              <Input placeholder="Please Enter Last Name" />
            </Form.Item>
            <Form.Item name="Email" label="Email" 
            rules={[
              {
                required: true,
                message: 'Please input Email',
              }
            ]}
            >
              <Input placeholder="Please Enter Email" />
            </Form.Item>
            <Form.Item name="PhoneNumber" label="Phone Number" 
            rules={[
              {
                required: true,
                message: 'Please input your Phone number',
              }
            ]}
            >
              <Input placeholder="Please Enter your Phone Number" />
            </Form.Item>

            <Form.Item name="Address" label="Address" 
            rules={[
              {
                required: true,
                message: 'Please input Address',
              }
            ]}
            >
              <Input placeholder="Please Enter Address" />
            </Form.Item>
            
            <div style={{textAlign: "right"}}>
            {userRole == 'Manager' && (
              <Button
                type="primary"
                loading={loading}
                htmlType="submit"
              >
                Save
              </Button>
          )}
            <Button type="default" htmlType="button" onClick={() => history('/customers')}>
              Back
            </Button>
              </div>
          </Form>
          </Col>
        </Row>
    </div>
  );
}
export default CustForm;