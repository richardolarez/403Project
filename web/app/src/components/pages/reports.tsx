import React, { useState } from 'react';
import { Row, Col, Button, Input } from 'antd';
import axios from 'axios';

const Reports = () => {
  const [result, setResult] = useState('');
  const [lastClickedButton, setLastClickedButton] = useState<number | null>(null);

  const handleButtonClick = async (buttonNumber: number) => {
    try {
      let response;

      switch (buttonNumber) {
        case 1:
          response = await axios.get(`http://localhost:8080/inventoryLogs`);
          break;
        case 2:
          response = await axios.get(`http://localhost:8080/financialLogs`);
          break;
        case 3:
          response = await axios.get(`http://localhost:8080/logs`);
          break;
        default:
          break;
      }

      if (response) {
        setResult(JSON.stringify(response.data));
        setLastClickedButton(buttonNumber);
      }
    } catch (error) {
      console.error(`Error fetching data for button ${buttonNumber}:`, error);
    }
  };

  return (
    <Row justify="center" align="top" style={{ height: '100vh' }}>
      <Col span={12}>
        <Row justify="center" gutter={[16, 16]} style={{ marginBottom: '10px' }}>
          <Col>
            <Button type="primary" onClick={() => handleButtonClick(1)}>
              Inventory Logs
            </Button>
          </Col>
          <Col>
            <Button type="primary" onClick={() => handleButtonClick(2)}>
              Fincancial Logs
            </Button>
          </Col>
          <Col>
            <Button type="primary" onClick={() => handleButtonClick(3)}>
              All Logs
            </Button>
          </Col>
        </Row>
        <Row justify="center">
          <Col span={24}>
            <Input.TextArea
              value={result}
              autoSize={{ minRows: 10, maxRows: 20 }}
              style={{ width: '100%', height: '300px' }}
              readOnly
            />
          </Col>
        </Row>
      </Col>
    </Row>
  );
};

export default Reports;
