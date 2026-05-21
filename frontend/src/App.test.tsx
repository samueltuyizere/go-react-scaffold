import { render, screen, fireEvent } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import App from "./App";

describe("App", () => {
  it("renders the heading", () => {
    render(<App />);
    expect(screen.getByText("Vite + React")).toBeInTheDocument();
  });

  it("increments count on button click", () => {
    render(<App />);
    const button = screen.getByRole("button", { name: /count is/i });
    expect(button).toHaveTextContent("count is 0");
    fireEvent.click(button);
    expect(button).toHaveTextContent("count is 1");
  });
});
