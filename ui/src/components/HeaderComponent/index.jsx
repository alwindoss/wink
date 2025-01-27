import React from "react";
import { Container, Nav, Navbar } from "react-bootstrap";

const HeaderComponent = () => {
  return (
    <div style={{ marginBottom: "10px" }}>
      <Navbar style={{ backgroundColor: "#404040" }}>
        <Container>
          <Navbar.Brand href="#home" style={{ color: "white" }}>
            Win the WEEK
          </Navbar.Brand>
          {/* <Nav className="me-auto">
            <Nav.Link href="#home" style={{ color: "white" }}>
              Home
            </Nav.Link>
            <Nav.Link href="#features" style={{ color: "white" }}>
              Features
            </Nav.Link>
            <Nav.Link href="#pricing" style={{ color: "white" }}>
              Pricing
            </Nav.Link>
          </Nav> */}
        </Container>
      </Navbar>
    </div>
  );
};

export default HeaderComponent;
