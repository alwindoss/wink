import { transitions, positions, Provider as AlertProvider } from "react-alert";
import AlertTemplate from "react-alert-template-basic";
import { Route, Routes, BrowserRouter as Router } from "react-router-dom";
import Home from "./pages/Home";
import NotFound from "./pages/NotFound";
import SomePage from "./pages/SomePage";
const options = {
  position: positions.TOP_RIGHT,
  timeout: 5000,
  offset: "30px",
  // you can also just use 'scale'
  transition: transitions.SCALE,
};

function App() {
  return (
    <AlertProvider template={AlertTemplate} {...options}>
        <Router>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/somepage" element={<SomePage />} />
            <Route path="*" element={<NotFound />} />
            <Route path="/notfound" element={<NotFound />} />
          </Routes>
        </Router>
    </AlertProvider>
  );
}

export default App;
