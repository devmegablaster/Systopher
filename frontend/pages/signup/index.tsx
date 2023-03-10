import React, { useEffect } from "react";
import { Alert } from "@mantine/core";
import { useSelector } from "react-redux";
import { IconAlertCircle } from "@tabler/icons";
import Router from "next/router";
import SignUpForm from "../../components/SignUpPage/SignUpForm";
import Error from "../../components/Utils/Notifications/Error";

export default function signup() {
  const [error, setError] = React.useState("");
  const [showError, setShowError] = React.useState(false);

  const signUpStatus = useSelector((state: any) => state.auth.signedUp);

  console.log(signUpStatus);

  useEffect(() => {
    if (signUpStatus) {
      Router.push("/signup/verify");
    }
  }, [signUpStatus]);

  return (
    <div className="bg-gray-900 h-screen w-screen scrollbar-hide overflow-x-hidden flex justify-between">
      <Error
        heading="Sign Up Error!"
        error={error}
        setShowError={setShowError}
        showError={showError}
      />
      <div className="md:w-1/2 bg-gray-900 w-full scrollbar-hide h-full overflow-y-scroll flex items-center justify-center">
        <SignUpForm
          setShowError={setShowError}
          showError={showError}
          error={error}
          setError={setError}
        />
      </div>
      <div className="w-1/2 bg-gray-900 md:flex hidden h-full items-center justify-center">
        <img src="login.svg" alt="" />
      </div>
    </div>
  );
}
