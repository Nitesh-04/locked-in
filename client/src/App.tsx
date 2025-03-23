import { SignIn, SignOutButton, useUser } from "@clerk/clerk-react";

export default function App() {
  const { isSignedIn, user } = useUser();

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      {isSignedIn ? (
        <>
          <h1>Welcome, {user?.fullName}!</h1>
          <SignOutButton />
        </>
      ) : (
        <>
          <p>Hello</p>
          <SignIn />
        </>
      )}
    </div>
  );
}
