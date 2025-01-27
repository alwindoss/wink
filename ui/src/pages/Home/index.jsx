import React from "react";
import { Card, Col, Container, Row } from "react-bootstrap";
import CardComponent from "../../components/CardComponent";
import HeadingComponent from "../../components/HeadingComponent";

const Home = () => {
  return (
    <div>
      <Container>
        <Row className="justify-content-md-center">
          <Col
            style={{
              border: "2px solid #404040",
              margin: "10px",
              borderRadius: "5px",
            }}
          >
            <HeadingComponent heading="Goals" />
            <CardComponent />
            <CardComponent />
            <CardComponent />
            <CardComponent />
            <CardComponent />
            <CardComponent />
          </Col>
          <Col
            style={{
              border: "2px solid #404040",
              margin: "10px",
              borderRadius: "5px",
            }}
          >
            <HeadingComponent heading="Accomplishments" />
            <CardComponent />
            <CardComponent />
            <CardComponent />
            <CardComponent />
            <CardComponent />
            <CardComponent />
          </Col>
        </Row>
      </Container>
    </div>
  );
};

export default Home;
