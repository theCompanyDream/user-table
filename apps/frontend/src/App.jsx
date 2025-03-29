import React, { Suspense, useState, useMemo} from "react";
import { Route, Routes } from "react-router-dom";

import { ClockLoader  } from "react-spinners"


import {UserContext} from "./components";

// Error Page
const Page404 = React.lazy(() => import('./components/pages/notFound'));

const App = () => {
  const defaultBookContext = {
    users: [],
  };

  const [users, setUsers] = useState(defaultBookContext);
  const contextMemo = useMemo(
    () => ({ users, setUsers }),
    [users]
  );

  return (
    <UserContext.Provider value={contextMemo}>
      <Layout>
        <Suspense fallback={<ClockLoader  color="#FFF200" size={50} />}>
          <Routes>
            <Route index path="/" element={<List />} />
            <Route path="/about" element={<About />} />
            <Route path="*" element={<Page404 />} />
          </Routes>
        </Suspense>
      </Layout>
    </UserContext.Provider>
  );
};

export default App;