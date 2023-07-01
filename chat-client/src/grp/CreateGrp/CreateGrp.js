import React, { useState, useEffect } from "react";
import "./CreateGrp.css";
import GrpIcon from "../GrpIcon/GrpIcon";
import InputWithIcon from "../../InputWithIcon/InputWithIcon";
import Search from "../../Search/Search";
import Chat from "../../chat/Chat/Chat";
import GoButton from "../../Auth/GoButton/GoButton";
import Member from "../Member/Member";

const CreateGrp = () => {
  const [members, setMembers] = React.useState([]);
  const [groupName, setGroupName] = React.useState("");
  const [selectedMembers, setSelectedMembers] = React.useState([]);
  const [searchResults, setSearchResults] = React.useState([]);

  const getMembers = async () => {
    // const response = await fetch('https://example.com/api/users');
    // const data = await response.json();
    let data = {
      status: "success",
      message: "users List retrieved successfully",
      data: [
        {
          user_id: "123",
          username: "user",
        },
        {
          user_id: "456",
          username: "user1",
        },
        {
          user_id: "789",
          username: "user2",
        },
        {
          user_id: "91",
          username: "user3",
        },
        {
          user_id: "434",
          username: "user4",
        },
        {
          user_id: "438",
          username: "user 1",
        },
      ],
    };

    if (data.status === "success") {
      setMembers(data.data);
    }
  };

  useEffect(() => {
    getMembers();
  }, []);

  const handleMemberClick = (user_id) => {
    const newMembers = [...selectedMembers, user_id];
    setSelectedMembers(newMembers);
  };
  
  const handleSearch = (searchTerm) => {
    if (searchTerm === "") {
      // If search term is empty, reset the search results to show all chats
      setSearchResults([]);
    } else {
      // Filter the chats based on the search term
      const results = members.filter((member) =>
        member.username.toLowerCase().includes(searchTerm.toLowerCase())
      );
      setSearchResults(results);
      console.log(results);
    }
  };

  const handleSubmit = () => {
    const postData = {
      group_name: groupName,
      members: selectedMembers,
    };
    console.log(postData)
    setSelectedMembers([]);
    // make POST request to API
    // const url = " /api/group/create";
    // const options = {
    //   method: "POST",
    //   headers: {
    //     "Content-Type": "application/json",
    //   },
    //   body: JSON.stringify(postData),
    // };
  
    // fetch(url, options)
    //   .then((response) => {
    //     if (response.status === 200) {
    //       console.log("Group created successfully!");
    //     } else {
    //       console.log("Error creating group!");
    //     }
    //   })
    //   .catch((error) => {
    //     console.log(error);
    //   });
  };
  

  return (
    <div className="createGrp">
      <div className="createGrp__grpIcon">
        <GrpIcon />
      </div>
      <div className="createGrp__Inp">
        <InputWithIcon
          placeholderText="Group name"
          icon="fa fa-users icon"
          onChange={(e) => setGroupName(e.target.value)}
        />
      </div>
      <div className="createGrp__addpartic">
        <h3 className="createGrp__addpartic--txt">Add Participants</h3>
      </div>
      <div className="createGrp__search">
        <Search onSearch={handleSearch} />
      </div>
      <div className="createGrp__members">
        {searchResults.length > 0
          ? searchResults.map((member) => (
            <Member
            key={member.user_id}
            name={member.username}
            id={member.user_id}
            onClick={handleMemberClick}
          />
          ))
        : members.map((member) => (
          <Member
            key={member.user_id}
            name={member.username}
            id={member.user_id}
            onClick={handleMemberClick}
          />
        ))}
      </div>
      <div className="createGrp__go">
        <GoButton handleSubmit={handleSubmit} />
      </div>
    </div>
  );
};

export default CreateGrp;
