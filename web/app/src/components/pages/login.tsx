import React from 'react';
import { Form, Input, Button, Row, Col, message, Modal } from 'antd';
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
  const [isModalVisible, setIsModalVisible] = React.useState(false);
  const [newPassword, setNewPassword] = React.useState('');
  const [currentUsername, setCurrentUsername] = React.useState('');
  const [currentPassword, setCurrentPassword] = React.useState('');
  const [shouldUpdatePassword, setShouldUpdatePassword] = React.useState(false);
  const [resolveModalPromise, setResolveModalPromise] = React.useState<(() => void) | null>(null);


  const handleOk = async () => {
    if (newPassword && currentUsername) {
      try {
        await updatePassword(currentUsername, currentPassword, newPassword);
        if (resolveModalPromise) {
          resolveModalPromise(); 
        }
        setIsModalVisible(false);
        // Reset the states after successful password update
        setCurrentUsername('');
        setNewPassword('');
        reattemptLogin(currentUsername, newPassword);
      } catch (error) {
        console.error('Error updating password:', error);
      }
    }
  };

  const reattemptLogin = async (username: string, newPassword: string) =>     
  fetch('http://localhost:8080/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ username, newPassword })
  })
  .then(response => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
  })
  .then(async (employee) => {
    if(employee.RequiresNewPass) {
      setCurrentUsername(username);
      setCurrentPassword(newPassword);
      await showModal();
    } 
    console.log('Authenticated employee:', employee);
    sessionStorage.setItem('authenticated', 'true');
    sessionStorage.setItem('UserFName', employee.FirstName);
    sessionStorage.setItem('UserRole', employee.Role)
    window.location.reload();
  });


  // useEffect to handle the actual password update process
  React.useEffect(() => {
    const performPasswordUpdate = async () => {
      if (shouldUpdatePassword && newPassword && currentUsername) {
        try {
          await updatePassword(currentUsername,currentPassword, newPassword);
          // Reset states after successful password update
          setCurrentUsername('');
          setNewPassword('');
          setShouldUpdatePassword(false);
        } catch (error) {
          console.error('Error updating password:', error);
        }
      }
    };

    performPasswordUpdate();
  }, [shouldUpdatePassword, newPassword, currentUsername]);

  const showModal = () => {
    setIsModalVisible(true);
    return new Promise<void>(resolve => {
      setResolveModalPromise(() => resolve);
    });
  };
  
  const handleCancel = () => {
    setIsModalVisible(false);
    setNewPassword('');
  };

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setNewPassword(e.target.value);
  };

  const updatePassword = async (username: string, oldPassword: string, newPassword: string) => {
    //Data package
    const data = {
      username: username,
      oldPassword: oldPassword,
      newPassword: newPassword
    };

    // API call to update pass
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
    } catch (error) {
      console.error('Error:', error);
      throw error; 
    }
  };

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
    .then(async (employee) => {
      if(employee.RequiresNewPass) {
        setCurrentUsername(values.username);
        setCurrentPassword(values.password);
        await showModal();
      } 
      console.log('Authenticated employee:', employee);
      sessionStorage.setItem('authenticated', 'true');
      sessionStorage.setItem('UserFName', employee.FirstName);
      sessionStorage.setItem('UserRole', employee.Role)
      window.location.reload();
    })
    .catch(error => {
      if(loginAttempts >= 5){

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
    <Modal
    title="Change Password"
    open={isModalVisible}
    onOk={handleOk}
    onCancel={handleCancel}
  >
    <Input.Password 
      value={newPassword} 
      onChange={handlePasswordChange} 
      placeholder="Enter new password" 
    />
  </Modal>  
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