import React from 'react';
import { Form, Input, Button, Row, Col, message } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import Password from 'antd/es/input/Password';

const Login: React.FC = () => {
  const [loginAttempts, setLoginAttempts] = React.useState(0);
  const [accountLocked, setAccountLocked] = React.useState(false);
  const [userLoginAttempts, setUserLoginAttempts] = React.useState<{ [username: string]: number }>({});
  //local storage for persistence. shadowRealm = locked users.
  const shadowRealm = React.useRef<string[]>(JSON.parse(localStorage.getItem('shadowRealm') || '[]'));
  const barrierToEntry = React.useRef<string[]>(JSON.parse(localStorage.getItem('barrierToEntry') || '[]'));
  //const passAlreadyChanged = (username: string) => barrierToEntry.current.includes(username); 
  const isAccountLocked = (username: string) => shadowRealm.current.includes(username);

  console.log(barrierToEntry);

  const onFinish = (values: { username: string, password: string , loginAttempts: number, accountLocked: boolean}) => {
    //Check if user is banished first
    if(isAccountLocked(values.username)){
      message.error('Account locked. Please contact an administrator');
      return;
    }

    //const isFirstLogin = !passAlreadyChanged(values.username);

    const updatePassword = async (username: string, oldPassword: string) => {
      return new Promise<void>(async (resolve, reject) => {
        let newPassword = prompt("Please enter a new password to continue: ");
        
        if (newPassword != null && newPassword.length > 6 && newPassword !== "") {
          const data = {
            username: username,
            oldPassword: oldPassword,
            newPassword: newPassword
          };
    
          try {
            const response = await fetch('http://localhost:8080/updatePassword', {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json',
              },
              body: JSON.stringify(data),
            });
    
            if (!response.ok) {
              throw new Error(`HTTP error! status: ${response.status}`);
            }
    
            //barrierToEntry.current.push(username);
            //localStorage.setItem('barrierToEntry', JSON.stringify(barrierToEntry.current));
            resolve();
          } catch (error) {
            console.error('Error:', error);
            reject(error);
          }
        } else {
          await updatePassword(username, oldPassword);
        }
      });
    };

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
    .then(async (employee) => {

      //Users must change password on FIRST login
      //passAlreadyChanged = list of users who have alreadyb changed their password
      //Calls updatePassword which takes current username and password as params
      // if (!passAlreadyChanged(values.username)) {
      //   await updatePassword(values.username, values.password);
      // }

      //Crazy how just doing it properly the first time is the easiest/fastest solution.
      if(employee.RequiresNewPass){
        updatePassword(values.username, values.password);
      }

      console.log('Authenticated employee:', employee);
      sessionStorage.setItem('authenticated', 'true');
      sessionStorage.setItem('UserFName', employee.FirstName);
      sessionStorage.setItem('UserRole', employee.Role);
      window.location.reload();
      
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