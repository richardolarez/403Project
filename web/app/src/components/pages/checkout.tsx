import React from 'react';
import { Button, Input } from 'antd';


const Checkout = () => {
      return (
        <div>
          <div style={{ display: 'flex', flexDirection: 'column' }}>
            <Input placeholder="Customer ID" style={{ width: 200, marginRight: 10 }} />
            <Input placeholder="ItemID" style={{ width: 200, marginRight: 10 }} />
            <Input placeholder="Qty" style={{ width: 100, marginRight: 10 }} />
            <Button style={{ width: 100, marginRight: 10 }}>+</Button>
            <Button type="primary" style={{ width: 200, marginRight: 10 }}>Add to Receipt</Button>
          </div>
        </div>
      );
    };
export default Checkout;