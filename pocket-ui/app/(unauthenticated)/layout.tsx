import type { Metadata } from "next";
import { cookies } from 'next/headers';

import { Navbar } from "./navbar";

export const metadata: Metadata = {
  title: "aimtrainer | sharpen your aim",
  description: "...",
};

export default async function UnauthenticatedLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const cookieStore = await cookies();
  const theme = cookieStore.get('theme')?.value || "light";

  return (
    <>
      <Navbar theme={theme} />
      <div>
        <section>
          {children}
        </section>
      </div>
    </>
  );
}
