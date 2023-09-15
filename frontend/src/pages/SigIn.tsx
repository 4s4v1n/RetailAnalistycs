import { Button, Card, Label, TextInput } from "flowbite-react";
import { useState, type FC, useContext } from "react";
import { Context } from "..";
import { observer } from "mobx-react-lite";
import { useNavigate } from "react-router-dom";
import { HOME_ROUTE, SIGIN_ROUTE } from "../utils/const";
import { notifyError, notifySucces } from "../components/Notify/Notify";


const SigIn: FC = () => {
  const [login, setlogin] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const { userStore } = useContext(Context)
  const navigate = useNavigate();

  const handleClick = async (event: React.MouseEvent<HTMLElement>) => {
    event.preventDefault()
    const err = await userStore.login(login, password)
    if (err.code === 200) {
      notifySucces("success login")
      navigate(HOME_ROUTE)
    } else {
      notifyError("invalid auth");
      navigate(SIGIN_ROUTE);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center px-6 lg:h-screen lg:gap-y-12">
      <Card
        horizontal
        imgSrc="/img/auth/login.png"
        imgAlt=""
        className="w-full md:max-w-screen-sm [&>img]:hidden md:[&>img]:w-60 md:[&>img]:p-0 md:[&>*]:w-full md:[&>*]:p-16 lg:[&>img]:block"
      >
        <h1 className="mb-3 text-2xl font-bold dark:text-white md:text-3xl">
          Sign in to platform
        </h1>
        <form>
          <div className="mb-4 flex flex-col gap-y-3">
            <Label htmlFor="login">Your login</Label>
            <TextInput
              id="login"
              name="login"
              placeholder="user"
              type="login"
              onChange={e => setlogin(e.target.value)}
              value={login}
            />
          </div>
          <div className="mb-6 flex flex-col gap-y-3">
            <Label htmlFor="password">Your password</Label>
            <TextInput
              id="password"
              name="password"
              placeholder="••••••••"
              type="password"
              onChange={e => setPassword(e.target.value)}
              value={password}
            />
          </div>
          <div className="mb-6">
            <Button
              onClick={handleClick}
              color="primary"
              type="Submit"
              className="w-full lg:w-auto">
              Login to your account
            </Button>
          </div>
        </form>
      </Card>
    </div>
  );
};

export default observer(SigIn);