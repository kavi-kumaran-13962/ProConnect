import React, { useState } from "react";
import "./HomeLayout.css";
import Nav from "../Nav/Nav";
import LogoHeader from "../LogoHeader/LogoHeader";
import Search from "../chat/Search/Search";
import Chat from "../chat/Chat/Chat";
import Profile from "../chat/Profile/Profile";
import ChatHeader from "../chat/ChatHeader/ChatHeader";
import MsgRec from "../chat/MsgRec/MsgRec";
import MsgSent from "../chat/MsgSent/MsgSent"
import ChatList from "../chat/ChatList/ChatList";
import { Outlet,Route,Routes } from "react-router-dom";
import MsgList from "../chat/MsgList/MsgList";
import GrpMsgList from "../chat/GrpMsgList/GrpMsgList";
import MsgInput from "../chat/MsgInput/MsgInput";
export default function HomeLayout() {
  return (
    <div className="homelayout">
      <div className="homelayout__logoheader">
        <LogoHeader />
      </div>
      <div className="homelayout__nav">
        <Nav />
      </div>
      <div className="homelayout__body">
        <div className="homelayout__body--profile">
          <Profile/>
        </div>
        <div className="homelayout__body--search">
          <Search />
        </div>
        <div className="homelayout__body--chatList">
        <Outlet/>
        </div>
      </div>
      <div className="homelayout__chat">
        <Routes>
            <Route path="chats/:chatId" element={<MsgList />} />
            <Route path="grpchats/:chatId" element={<GrpMsgList />} />
          </Routes>
        </div>
        <div className="homelayout__msgInp">
            <MsgInput/>
        </div>
      </div>
  );
}
