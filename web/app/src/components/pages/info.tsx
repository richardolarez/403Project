//src/components/pages/list.tsx
import React, {useEffect, useState} from 'react';
import {Table, Row, Col, Button, Typography} from 'antd';
import {useNavigate} from 'react-router';
import axios from 'axios';

const {Title} = Typography;

const Info = () => {
    const history = useNavigate();
    const [allData, setAllData] = useState([]);
  
    useEffect(() => {
      axios.get(`http://localhost:8080/pharmacies`).then(res => {
        setAllData(res.data);
      });
    }, []);
  
    const information = [
      {
    Name: "Pharmacy A",
    Location: "Location A",
    Currently: true,
    website: "pharmacyA.com",
    owner: "Owner A",
    phoneNumber: 1234567890,
    Hours: "9:00 AM - 5:00 PM"
      },
    ];

    const data: { name: any; location: any; isOpen: any; website: any; owner: any; phoneNumber: any, workingHours: any}[] = [];
    allData.map((pharmInfo: any) => {
        data.push({
          name: pharmInfo.name,
          location: pharmInfo.location,
          isOpen: pharmInfo.isOpen,
          website: pharmInfo.website,
          owner: pharmInfo.owner,
          phoneNumber: pharmInfo.phoneNumber,
          workingHours: pharmInfo.workingHours,
      })
      return data;
    });
  
  console.log(data)

  return (
      <div>
        
          <h1>Pharmacy Name</h1>
          <h3>Location</h3>
          <h3>Website</h3>
          <h3>Owner</h3>
          <h3>Phone Number</h3>
          <h3>Operating Hours</h3>
          
      </div>
      
    );
  }
export default Info;