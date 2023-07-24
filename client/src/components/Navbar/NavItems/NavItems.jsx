import React from "react";
import { classNames } from "../../../lib/utils/utils";
import { Link } from "react-router-dom";
import "../Navbar.css";

export default function NavItems({ navigation }) {
  return (
    <div className="flex flex-1 items-center justify-center sm:items-stretch sm:justify-start">
      <div className="flex flex-shrink-0 items-center">
        <img
          className="block h-10 w-10 lg:hidden"
          src="https://media.licdn.com/dms/image/C5603AQFbUePI_2E0Cg/profile-displayphoto-shrink_400_400/0/1616816921606?e=1695859200&v=beta&t=7ZAXnYgPeT3tTLk39Bs0NBAEbnyc2B7oL4cPx8K45TU"
          alt=""
        />
        <img
          className="hidden h-10 w-10  lg:block"
          src="https://media.licdn.com/dms/image/C5603AQFbUePI_2E0Cg/profile-displayphoto-shrink_400_400/0/1616816921606?e=1695859200&v=beta&t=7ZAXnYgPeT3tTLk39Bs0NBAEbnyc2B7oL4cPx8K45TU"
          alt=""
        />
      </div>
      <div className="hidden sm:ml-6 sm:block">
        <div className="flex items-center justify-center space-x-4 h-full">
          {navigation.map((item) => (
            <Link
              to={item.href}
              key={item.name}
              className={classNames(
                "relative inline-block",
                item.current ? "text-black" : "text-gray-400",
                "custom-link"
              )}
              aria-current={item.current ? "page" : undefined}
            >
              {item.name}
            </Link>
          ))}
        </div>
      </div>
    </div>
  );
}
