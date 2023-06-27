import React, { useState, useEffect } from "react";
import Chat from "../Chat/Chat";
import Search from "../../Search/Search";
import "./ChatList.css"

const ChatList = () => {
  const [chats, setChats] = useState([]);
  const [searchResults, setSearchResults] = useState([]);
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
          user_id: "100",
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
          user_id: "101",
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
          user_id: "102",
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
          user_id: "103",
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
          user_id: "104",
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
          user_id: "105",
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
          user_id: "106",
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
          user_id: "107",
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

  const handleSearch = (searchTerm) => {
    if (searchTerm === "") {
      // If search term is empty, reset the search results to show all chats
      setSearchResults([]);
    } else {
      // Filter the chats based on the search term
      const results = chats.filter((chat) =>
        chat.username.toLowerCase().includes(searchTerm.toLowerCase())
      );
      setSearchResults(results);
      console.log(results);
    }
  };

  return (
    <div className="chatList">
      <div className="chatList__search">
        <Search onSearch={handleSearch} />
      </div>
      <div className="chatList__chats">
        {searchResults.length > 0
          ? searchResults.map((chat) => (
              <Chat
                key={chat.user_id}
                id={chat.user_id}
                name={chat.username}
                lastMessage={chat.last_message.content}
              />
            ))
          : chats.map((chat) => (
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