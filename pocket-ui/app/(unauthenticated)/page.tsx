export default function IndexPage() {
  return (
    <>
      <div className="h-[540px] bg-blue-400 py-3">
        <div className="h-full max-w-4xl mx-auto cursor-pointer hidden">
          <div className="flex justify-center items-center h-full">
            <h2 className="text-white text-5xl font-medium">Click to train</h2>
          </div>
        </div>
        <div className="h-full max-w-4xl mx-auto cursor-pointer">
          <div className="flex justify-center items-center h-full">
            <h2 className="text-white text-5xl font-medium">Hello</h2>
          </div>
        </div>
      </div>
      <div className="container mx-auto">
        <div>Hello World</div>
        <div>this is page tsx</div>
      </div>
    </>
  );
}
