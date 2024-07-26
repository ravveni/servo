// servo -- the perfect tool for multi-device ux/ui testing over LAN
// Copyright (C) 2024 ravveni

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	port := "8080"

	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	fs := http.FileServer(http.Dir(currentDir))
	http.Handle("/", fs)

	fmt.Println("Serving", currentDir, "on port", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
