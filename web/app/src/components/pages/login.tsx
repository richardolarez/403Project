import React from 'react';
import { Form, Input, Button, Row, Col, message } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';

const Login: React.FC = () => {
  const [loginAttempts, setLoginAttempts] = React.useState(0);
  const [accountLocked, setAccountLocked] = React.useState(false);
  const [userLoginAttempts, setUserLoginAttempts] = React.useState<{ [username: string]: number }>({});
  //local storage for persistence. shadowRealm = locked users.
  const shadowRealm = React.useRef<string[]>(JSON.parse(localStorage.getItem('shadowRealm') || '[]'));
  const isAccountLocked = (username: string) => shadowRealm.current.includes(username);
  const onFinish = (values: { username: string, password: string , loginAttempts: number, accountLocked: boolean}) => {
    //Check if user is banished first
    if(isAccountLocked(values.username)){
      message.error('Account locked. Please contact an administrator');
      return;
    }
    // Call the /login endpoint to authenticate the employee
    fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ username: values.username, password: values.password })
    })
    .then(response => {
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      return response.json();
    })
    .then(employee => {
      if(!accountLocked){
        console.log('Authenticated employee:', employee);
        sessionStorage.setItem('authenticated', 'true');
        sessionStorage.setItem('UserFName', employee.FirstName);
        sessionStorage.setItem('UserRole', employee.Role)
        window.location.reload();
      } 
      // else{
      //   message.error('Your account has been locked. Please contact an administrator');
      // }
    })
    .catch(error => {
      if(loginAttempts >= 5){
        //setAccountLocked(true); -- deprecated

        //Send them to the shadow realm aka lock their account
        shadowRealm.current.push(values.username);     
        localStorage.setItem('shadowRealm', JSON.stringify(shadowRealm.current));
        message.error('Account locked. Please contact an administrator');

        //Reset counter for next login attempts.
        setLoginAttempts(loginAttempts => loginAttempts = 0);
      } else{
        console.error('Failed to authenticate employee:', error);
        // TODO: Handle authentication error
        message.error('Incorrect credentials. Please check username and password');
        setLoginAttempts(loginAttempts => loginAttempts +=1);
      }
    });
  };

  return (
  <div>  
    
  <Row justify="center" align="middle" style={{ minHeight: '100vh' }}>
      <Col span={8}>
        <img src="/poos_logo_full.png" alt="Logo" style={{width: '100%', marginBottom: '10%'}}/>
          <Form
          name="login"
          initialValues={{ remember: true }}
          onFinish={onFinish}
          style={{ background: 'white', padding: '20px', boxShadow: '0px 0px 10px rgba(0, 0, 0, 0.2)' }}
        >
          <Form.Item
            name="username" label="Username"
            rules={[{ required: true, message: 'Please enter your username!' }]}
          >
            <Input prefix={<UserOutlined />} placeholder="Username" />
          </Form.Item>

          <Form.Item
            name="password" label="Password"
            rules={[{ required: true, message: 'Please enter your password!' }]}
          >
            <Input.Password prefix={<LockOutlined />} placeholder="Password" />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" style={{ width: '100%' }}>
              Log In
            </Button>
          </Form.Item>
        </Form>
      </Col>
    </Row>
    </div>
  );
};

export default Login;