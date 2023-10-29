//src/components/pages/prescriptions.tsx
import React, {useState} from 'react';
import {Typography, message} from 'antd';
import axios from 'axios';
import {useNavigate} from 'react-router';
const {Title} = Typography;
const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};
const Prescriptions = () => {
  const [loading, setLoading] = useState(false);
  const history = useNavigate();
  
const handleSubmit = (values: any) => {
    setLoading(true);
    axios.post(`http://localhost:8080/fillPrescription`, 
      values
    )
    .then(res => {
      setLoading(false);
      message.success('Prescription Added Successfully!');
      history('/list');
    })
    .catch(error => {
      setLoading(false);
      message.error(error);
    })
  }
return (
    <div>
        TEST
    </div>
  );
}
export default Prescriptions;