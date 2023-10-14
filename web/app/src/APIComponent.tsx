import React, { useState, useEffect } from 'react';

const ApiComponent: React.FC = () => {
    const [items, setItems] = useState<{ id: number; name: string }[]>([]);
    const [newItem, setNewItem] = useState({ name: '' });
    const [updateItem, setUpdateItem] = useState({ name: '' });
    const [selectedItemId, setSelectedItemId] = useState<number | null>(null);
    const [postResponse, setPostResponse] = useState<number | null>(null); 
    const [patchResponse, setPatchResponse] = useState<number | null>(null); 

    const employeeApi = async (http: string, target: string, payload: JSON) => {
        try {
            let response: Response;

            if (http === "GET") {
                response = await fetch('http://localhost:8080/403proj/internal/service/employee', {
                method: 'GET',
                headers: {
                'Content-Type': 'application/json',
                },
                body: target,
            });

            const data = await response.json();
            setItems(data);
            }

            if (http === "POST") {
                response = await fetch('http://localhost:8080/403proj/internal/service/employee', {
                method: 'POST',
                headers: {
                'Content-Type': 'application/json',
                },
                body: JSON.stringify(newItem),
            });
            setPostResponse(response.status); 
            setNewItem({ name: '' });
            }

            if (http === "PATCH") {
                if (Object.keys(payload)[0] === null) return;
                    response = await fetch(`http://localhost:8080/403proj/internal/service/employee`, {
                    method: 'PATCH',
                    headers: {
                    'Content-Type': 'application/json',
                    },
                    body: target + JSON.stringify(payload),
                });
            setPatchResponse(response.status); 
            }
        } catch (error) {
        console.error('Employee API Error - ' + http + payload, error);
        }
    };

    return (
        <div>
          {/* Display POST response status */}
          {postResponse && <div id='post-response'>POST Response: {postResponse}</div>}
      
          {/* Display PATCH response status */}
          {patchResponse && <div id='patch-response'>PATCH Response: {patchResponse}</div>}
      
          {/* Display GET response status and items */}
          {items.length > 0 && (
            <div id='get-response'>
              {/* Display GET response status */}
              <div>GET Response: 200</div> {/* Assuming a successful GET request */}
      
              {/* Display GET response items */}
              <div>
                <h2>Items:</h2>
                <ul>
                  {items.map((item) => (
                    <li key={item.id}>{item.name}</li>
                  ))}
                </ul>
              </div>
            </div>
          )}
        </div>
      );
};

export default ApiComponent;

