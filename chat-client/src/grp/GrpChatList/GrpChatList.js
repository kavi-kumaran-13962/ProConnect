import React from "react";
import Chat from "../../chat/Chat/Chat";
import { useState, useEffect } from "react";
import AddBtn from "../AddBtn/AddBtn";
import Search from "../../chat/Search/Search";
import "./GrpChatList.css";
const GrpChatList = () => {
  const [chats, setChats] = useState([]);

  useEffect(() => {
    // Fetch data from the API
    // fetch('your-api-endpoint')
    //   .then(response => response.json())
    //   .then(data => setChats(data));
    let data = {
      status: "success",
      message: "Group chats retrieved successfully",
      data: [
        {
          group_id: "123",
          group_name: "Group 1",
          last_message: {
            message_id: "456",
            sender_id: "789",
            content: "Hello everyone!",
            timestamp: "2023-06-13T15:30:00Z",
          },
        },
        {
          group_id: "456",
          group_name: "Group 2",
          last_message: {
            message_id: "789",
            sender_id: "123",
            content: "Hey there!",
            timestamp: "2023-06-12T10:45:00Z",
          },
        },
        {
          group_id: "123",
          group_name: "Group 1",
          last_message: {
            message_id: "456",
            sender_id: "789",
            content: "Hello everyone!",
            timestamp: "2023-06-13T15:30:00Z",
          },
        },
        {
          group_id: "456",
          group_name: "Group 2",
          last_message: {
            message_id: "789",
            sender_id: "123",
            content: "Hey there!",
            timestamp: "2023-06-12T10:45:00Z",
          },
        },
        {
          group_id: "123",
          group_name: "Group 1",
          last_message: {
            message_id: "456",
            sender_id: "789",
            content: "Hello everyone!",
            timestamp: "2023-06-13T15:30:00Z",
          },
        },
        {
          group_id: "456",
          group_name: "Group 2",
          last_message: {
            message_id: "789",
            sender_id: "123",
            content: "Hey there!",
            timestamp: "2023-06-12T10:45:00Z",
          },
        },
        {
          group_id: "123",
          group_name: "Group 1",
          last_message: {
            message_id: "456",
            sender_id: "789",
            content: "Hello everyone!",
            timestamp: "2023-06-13T15:30:00Z",
          },
        },
        {
          group_id: "456",
          group_name: "Group 2",
          last_message: {
            message_id: "789",
            sender_id: "123",
            content: "Hey there!",
            timestamp: "2023-06-12T10:45:00Z",
          },
        },
        {
          group_id: "123",
          group_name: "Group 1",
          last_message: {
            message_id: "456",
            sender_id: "789",
            content: "Hello everyone!",
            timestamp: "2023-06-13T15:30:00Z",
          },
        },
        {
          group_id: "456",
          group_name: "Group 2",
          last_message: {
            message_id: "789",
            sender_id: "123",
            content: "Hey there!",
            timestamp: "2023-06-12T10:45:00Z",
          },
        },
        
      ],
    };

    setChats(data.data);
  }, []);

  return (
    <div className="grpChatList">
      <div className="grpChatList__search">
        <Search />
      </div>
      <div className="grpChatList__chats">
        {chats.map((chat) => (
          <Chat
            key={chat.group_id}
            id={chat.group_id}
            name={chat.group_name}
            lastMessage={chat.last_message.content}
          />
        ))}
      </div>
      <div className="grpChatList__grpCreateBtn">
        <AddBtn />
      </div>
    </div>
  );
};

export default GrpChatList;
