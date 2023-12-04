import React from 'react';
import { Card, Row, Col, Typography } from 'antd';

const { Title, Text } = Typography;

const Info = () => {
  const data = [
    {
      name: 'POOS',
      location: 'Location 1',
      isOpen: true,
      website: 'www.poos.com',
      owner: 'Owner 1',
      phoneNumber: '123-456-7890',
      workingHours: '9am - 5pm',
    },
  ];

  return (
    <div>
      <Row gutter={[40, 0]}>
        <Col span={24}>
          <Title level={2}>Pharmacy Information</Title>
        </Col>
      </Row>
      <Row gutter={[40, 20]}>
        {data.map((pharmacy, index) => (
          <Col key={index} xs={24} sm={12} md={8} lg={6}>
            <Card title={pharmacy.name}>
              <Text strong>Location: </Text>
              <Text>{pharmacy.location}</Text>
              <br />
              <Text strong>Currently Open: </Text>
              <Text>{pharmacy.isOpen ? 'Yes' : 'No'}</Text>
              <br />
              <Text strong>Website: </Text>
              <Text>{pharmacy.website}</Text>
              <br />
              <Text strong>Owner: </Text>
              <Text>{pharmacy.owner}</Text>
              <br />
              <Text strong>Phone Number: </Text>
              <Text>{pharmacy.phoneNumber}</Text>
              <br />
              <Text strong>Working Hours: </Text>
              <Text>{pharmacy.workingHours}</Text>
            </Card>
          </Col>
        ))}
      </Row>
    </div>
  );
}
export default Info;