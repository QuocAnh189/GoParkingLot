import { Suspense, lazy } from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Layout from "./layouts";

//page
const LoginPage = lazy(() => import("@pages/auth/SignIn"));
const RegisterPage = lazy(() => import("@pages/auth/SignUp"));
const DetectPage = lazy(() => import("@pages/detect/Detect"));
const ManagementCard = lazy(() => import("@pages/manage/ManageCard"));
const ManagementIO = lazy(() => import("@pages/manage/ManageIO"));

//component
import Loader from "@components/common/Loader";

function App() {
  return (
    <BrowserRouter>
      <Suspense fallback={<Loader />}>
        <Routes>
          <Route path="/" element={<LoginPage />} />
          <Route path="/signup" element={<RegisterPage />} />
          <Route path="/main" element={<Layout />}>
            <Route path="/main" element={<DetectPage />} />
            <Route path="/main/manage-card" element={<ManagementCard />} />
            <Route path="/main/manage-io" element={<ManagementIO />} />
            {/* <Route path="/main/setting" element={<RFID />} /> */}
          </Route>
        </Routes>
      </Suspense>
    </BrowserRouter>
  );
}

export default App;
