import React from 'react';
import MsgRec from '../../chat/MsgRec/MsgRec';
import MsgSent from '../../chat/MsgSent/MsgSent';
import ChatHeader from '../../chat/ChatHeader/ChatHeader';
import { useState, useEffect } from 'react';

import "./GrpMsgList.css"

const GrpMsgList = () => {
    const [data, setData] = useState({
      user_id: "",
      username: "",
      messages: []
    });
  
    useEffect(() => {
      // Fetch data from the API
      // fetch('your-api-endpoint')
      //   .then(response => response.json())
      //   .then(data => setData(data.data));
      let apiData = {
        "status": "success",
        "message": "Chat messages retrieved successfully",
        "data": {
          "user_id": "123",
          "username": "user1",
          "messages": [
            {
              "isSent": "true",
              "content": "Hello!",
              "timestamp": "2023-06-13T15:30:00Z"
            },
            {
              "isSent": "false",
              "content": "Hi there!",
              "timestamp": "2023-06-12T10:45:00Z"
            }
          ]
        }
      };    
      setData(apiData.data);
    }, []);
  
    return (
      <>
        <ChatHeader name={data.username} />
        <div className="msgList">
          {data.messages.map((message, index) => (
            message.isSent === "true" ? (
              <MsgSent key={index} message={message.content} />
            ) : (
              <MsgRec key={index} message={message.content} />
            )
          ))}
        </div>
      </>
    );
};

export default GrpMsgList;
