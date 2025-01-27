import React from "react";

const HeadingComponent = (props) => {
  return (
    <div
      style={{
        width: "100%",
        height: "3%",
        backgroundColor: "#404040",
        marginTop: "5px",
        marginBottom: "5px",
        borderRadius: "5px",
        textAlign: "center",
        color: "white",
        fontSize: "25px",
      }}
    >
      {props.heading}
    </div>
  );
};

export default HeadingComponent;
