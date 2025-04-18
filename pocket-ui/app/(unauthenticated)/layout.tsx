import type { Metadata } from "next";
import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';

import { auth } from "@/app/lib/auth";
import { Navbar } from "./navbar";

export const metadata: Metadata = {
  title: "e-pocket",
  description: "streamlined e-wallet experience",
};

export default async function UnauthenticatedLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const cookieStore = await cookies();
  const theme = cookieStore.get('theme')?.value || "light";
  const session = await auth();

  if (session != null) {
    return redirect("/wallet");
  }

  return (
    <div>
      {children}
    </div>
  );
}
