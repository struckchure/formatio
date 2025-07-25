import Footer from "@/layouts/baselayout/footer";
import Header from "@/layouts/baselayout/header";

export function ErrorPage() {
  return (
    <>
      <Header />
      <main className="p-4 text-center">
        <p>This page cannot be found.</p>
        <p className="mb-4">
          Check if you have added the route or the route link is correct
        </p>

        <p>
          you can edit this error page in{" "}
          <code className="bg-[#444] p-1 rounded-md">pages/error</code>
        </p>
      </main>
      <Footer />
    </>
  );
}
