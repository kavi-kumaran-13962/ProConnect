import React from 'react';
import "./CreateGrp.css"
import GrpIcon from "../GrpIcon/GrpIcon"
import InputWithIcon from "../../InputWithIcon/InputWithIcon"
const CreateGrp = () => {

  return (
    <>
        <GrpIcon/>
        <InputWithIcon placeholderText="Group name" icon="fa fa-users icon"/>
    </>
  );
};

export default CreateGrp;
