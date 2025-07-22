import { CiGrid41 } from "react-icons/ci";
import { Link, Navigate } from "react-router-dom";

import { Button } from "@/components/button";
import { useAuthStore } from "@/store";

export function HomePage() {
  const { isAuthenticated } = useAuthStore();
  if (isAuthenticated) return <Navigate to={"/dashboard"} />;

  return (
    <main>
      <section className="hero w-full my-4 flex-center">
        <div className="sm:w-[40rem] px-8 py-20 gap-8 grid center text-center">
          <h3 className="text-6xl font-semibold">
            Simplifying Devops for Developers
          </h3>
          <p className="text-xl">
            Streamline workflows, automate processes, and elevate productivity
            effortlessly.
          </p>
          <div className="flex-center gap-4">
            <Link to="/login">
              <Button>Get Started</Button>
            </Link>

            <Link to={"/register"}>
              <Button theme="outline">Learn More</Button>
            </Link>
          </div>
        </div>
      </section>

      <section className="py-10 flex-col-center gap-4">
        <h2 className="text-3xl text-center">
          A powerful and innovative feature set
        </h2>
        <p className="font-extralight text-xl">
          Built for developers, by developers
        </p>

        <div className="grid sm:grid-cols-3 gap-8 sm:px-16 px-6">
          {[...new Array(6).fill(0)].map(() => (
            <div className="relative p-6 border border-secondary shadow-[#2c7766] shadow-md rounded-lg">
              <h5 className="text-xl text-secondary font-medium">
                Faster Builds
              </h5>
              <p className="text-gray-300">
                Code builds 3x faster on Formatio than on other platforms via
                Formachs.
              </p>
              <div className="absolute bottom-0 right-0 px-4 py-3 text-secondary text-2xl border-secondary border-l border-t rounded-full">
                <CiGrid41 />
              </div>
            </div>
          ))}
        </div>
      </section>
    </main>
  );
}
