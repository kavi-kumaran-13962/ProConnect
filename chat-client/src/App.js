import './App.css';
import SignIn from './Auth/SignIn/SignIn'
import AuthInput from './InputWithIcon/InputWithIcon';
import MsgInput from './chat/MsgInput/MsgInput';
import MsgRec from './chat/MsgRec/MsgRec'
import MsgSent  from './chat/MsgSent/MsgSent';
import Chat from './chat/Chat/Chat';
import Search from './Search/Search';
import Nav from "./Nav/Nav";
import LogoHeader from "./LogoHeader/LogoHeader"
import AddBtn from './grp/AddBtn/AddBtn';
import HomeLayout from "./HomeLayout/HomeLayout"
import {Routes, Route} from "react-router-dom"
import ChatList from './chat/ChatList/ChatList';
import GrpChatList from './grp/GrpChatList/GrpChatList';
import MsgList from './chat/MsgList/MsgList';
import GrpMsgList from "./grp/GrpMsgList/GrpMsgList"
import CreateGrp from "./grp/CreateGrp/CreateGrp"
import Profile from './chat/Profile/Profile';
function App() {
  return (
    <Routes>
      <Route path="/" element={<HomeLayout/>}>
        <Route path="chats" element={<ChatList/>}>
          <Route path=":chatId" element={<MsgList/>} />
        </Route>
        <Route path="grpchats" element={<GrpChatList/>}>
          <Route path=":chatId" element={<GrpMsgList/>} />
        </Route>
        <Route path='createGrp' element={<CreateGrp/>}/>
        <Route path='profile' element={<Profile/>}/>
      </Route>
    </Routes>
  );
}


export default App;
