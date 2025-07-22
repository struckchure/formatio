import { FaCircle as DotIcon } from "react-icons/fa";
import {
  FiGithub as GithubIcon,
  FiInstagram as InstagramIcon,
  FiLinkedin as LinkedInIcon,
  FiTwitter as TwitterIcon,
} from "react-icons/fi";
import { Link } from "react-router-dom";

import { Button } from "@/components/button";
import formatioLogo from "@/images/formatio-logo-48-x-48.png";

function Footer() {
  return (
    <footer className="bg-[#333] text-center py-6">
      <div className="container mx-auto flex items-center justify-center gap-10">
        <div className="flex items-center justify-center gap-6">
          <Link
            to={"https://linkedin.com/company/formatio-cloud"}
            target="_blank"
          >
            <Button theme="ghost">
              <LinkedInIcon />
            </Button>
          </Link>

          <Link to={"https://twitter.com/formatio_cloud"} target="_blank">
            <Button theme="ghost">
              <TwitterIcon />
            </Button>
          </Link>

          <Link to={"https://instagram.com/formatio_cloud"} target="_blank">
            <Button theme="ghost">
              <InstagramIcon />
            </Button>
          </Link>

          <Link to={"https://github.com/overal-x"} target="_blank">
            <Button theme="ghost">
              <GithubIcon />
            </Button>
          </Link>
        </div>

        <span className="text-[4px]">
          <DotIcon />
        </span>

        <Link to="/terms">
          <Button theme="link">Terms</Button>
        </Link>

        <span className="text-[4px]">
          <DotIcon />
        </span>

        <Link to="/privacy">
          <Button theme="link">Privacy</Button>
        </Link>

        <span className="text-[4px]">
          <DotIcon />
        </span>

        <p className="flex items-center gap-2">
          <img src={formatioLogo} className="w-6" />

          <span>&copy; Formatio {new Date().getFullYear()}</span>
        </p>
      </div>
    </footer>
  );
}

export default Footer;
