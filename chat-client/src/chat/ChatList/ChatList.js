import React, { useState, useEffect } from "react";
import Chat from "../Chat/Chat";
import Search from "../Search/Search";
import "./ChatList.css"
const ChatList = () => {
  const [chats, setChats] = useState([]);

  useEffect(() => {
    // Fetch data from the API
    // fetch('your-api-endpoint')
    //   .then(response => response.json())
    //   .then(data => setChats(data));
    let data = {
      status: "success",
      message: "Chats retrieved successfully",
      data: [
        {
          user_id: "123",
          username: "user1",
          last_message: {
            message_id: "456",
            sender_id: "123",
            recipient_id: "789",
            content: "Hello!",
            timestamp: "2023-06-13T15:30:00Z",
          },
        },
        {
          user_id: "789",
          username: "user2",
          last_message: {
            message_id: "789",
            sender_id: "789",
            recipient_id: "123",
            content: "Hi there!",
            timestamp: "2023-06-12T10:45:00Z",
          },
        },
        {
          user_id: "123",
          username: "user1",
          last_message: {
            message_id: "456",
            sender_id: "123",
            recipient_id: "789",
            content: "Hello!",
            timestamp: "2023-06-13T15:30:00Z",
          },
        },
        {
          user_id: "789",
          username: "user2",
          last_message: {
            message_id: "789",
            sender_id: "789",
            recipient_id: "123",
            content: "Hi there!",
            timestamp: "2023-06-12T10:45:00Z",
          },
        },
        {
          user_id: "123",
          username: "user1",
          last_message: {
            message_id: "456",
            sender_id: "123",
            recipient_id: "789",
            content: "Hello!",
            timestamp: "2023-06-13T15:30:00Z",
          },
        },
        {
          user_id: "789",
          username: "user2",
          last_message: {
            message_id: "789",
            sender_id: "789",
            recipient_id: "123",
            content: "Hi there!",
            timestamp: "2023-06-12T10:45:00Z",
          },
        },
        {
          user_id: "123",
          username: "user1",
          last_message: {
            message_id: "456",
            sender_id: "123",
            recipient_id: "789",
            content: "Hello!",
            timestamp: "2023-06-13T15:30:00Z",
          },
        },
        {
          user_id: "789",
          username: "user2",
          last_message: {
            message_id: "789",
            sender_id: "789",
            recipient_id: "123",
            content: "Hi there!",
            timestamp: "2023-06-12T10:45:00Z",
          },
        },
      ],
    };

    setChats(data.data);
  }, []);

  return (
    <div className="chatList">
      <div className="chatList__search">
        <Search />
      </div>
      <div className="chatList__chats">
        {chats.map((chat) => (
          <Chat
            key={chat.user_id}
            id={chat.user_id}
            name={chat.username}
            lastMessage={chat.last_message.content}
          />
        ))}
      </div>
    </div>
  );
};

export default ChatList;
