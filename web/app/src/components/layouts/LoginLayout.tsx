// LoginLayout.tsx
import React, {ReactNode} from 'react';

interface LoginLayoutProps {
    children: ReactNode;
  }

const LoginLayout: React.FC<LoginLayoutProps> = ({ children }) => {
  return (
    <div>
      {children}
    </div>
  );
};

export default LoginLayout;