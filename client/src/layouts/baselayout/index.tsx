import { Outlet } from "react-router-dom";

import Footer from "@/layouts/baselayout/footer";
import Header from "@/layouts/baselayout/header";

export function BaseLayout() {
  return (
    <main>
      <Header />
      <Outlet />
      <Footer />
    </main>
  );
}
