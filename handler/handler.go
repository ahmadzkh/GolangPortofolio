/**
 * File handler.go
 *
 * Package handler provides HTTP request handling and template rendering for the web application.
 *
 *  * Dependencies:
 *
 * - `html/template`: Used for parsing and executing HTML templates.
 * - `log`: Used for logging errors and informational messages.
 * - `net/http`: Provides HTTP request and response handling.
 * - `path`: Assists in constructing file paths for template files.
 *
 * @ahmadzkh
 */
package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)

	if r.URL.Path != "/" {
		// Render 404 template
		tmpl, err := template.ParseFiles(
			path.Join("views", "header.html"),
			path.Join("views", "navbar.html"),
			path.Join("views", "404.html"),
			path.Join("views", "footer.html"),
		)
		if err != nil {
			log.Println("Error parsing templates:", err)
			http.Error(w, "Oops... internal server error", http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"title": "404 Not Found",
			"menu1": "About",
			"menu2": "Certified",
			"menu3": "Project",
			"menu4": "Contact",
		}

		err = tmpl.ExecuteTemplate(w, "404.html", data)
		if err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Oops... internal server error", http.StatusInternalServerError)
			return
		}
		return
	}

	// Parse all templates
	tmpl, err := template.ParseFiles(
		path.Join("views", "header.html"),
		path.Join("views", "navbar.html"),
		path.Join("views", "index.html"),
		path.Join("views", "intro.html"),
		path.Join("views", "certified.html"),
		path.Join("views", "projects.html"),
		path.Join("views", "contact.html"),
		path.Join("views", "extras.html"),
		path.Join("views", "footer.html"),
	)

	if err != nil {
		log.Println("Error parsing templates:", err)
		http.Error(w, "Oooppss.... error pages", http.StatusInternalServerError)
		return
	}

	// Data to be passed to template
	data := map[string]interface{}{
		"title":        "Azami",
		"menu1":        "About",
		"menu2":        "Certified",
		"menu3":        "Project",
		"menu4":        "Contact",
		"name":         "AHMAD ZAKY HUMAMI",
		"description":  "Saya adalah Web Developer yang berorientasi pada hasil dengan pengalaman lebih dari 3 tahun. Mahir dalam berbagai bahasa program dan berpengalaman dalam pengembangan web full-stack dengan memanfaatkan teknologi baru untuk mendorong inovasi.",
		"myPhone":      "+6285156200870",
		"imageProfile": "/assets/image/profiles.jpg",
		"vscodeLogo":   "/assets/image/visual-studio-code-1.svg",
		"phpLogo":      "/assets/image/php-icon.svg",
		"golangLogo":   "/assets/image/golang-1.svg",
		"pythonLogo":   "/assets/image/python-5.svg",
		"javaLogo":     "/assets/image/java-14.svg",
		"bpkLogo":      "https://jatim.bpk.go.id/wp-content/uploads/2020/08/cropped-Logo-BPK-icon-1.png",
		"stmikLogo":    "https://www.stmik.banisaleh.ac.id/view/logo/logo.png",
		"license1":     "Dicoding Indonesia",
		"certifLogo1":  "https://cdn.brandfetch.io/idK1RDMSRY/w/400/h/400/theme/dark/icon.jpeg?k=bfHSJFAPEG",
		"license2":     "Progate",
		"project1":     "/assets/image/landpage_pytricity.png",
		"project2":     "/assets/image/dashboard_inven.png",
		"project3":     "/assets/image/landpage_bkk.png",
		"project4":     "/assets/image/profile_instagramclone.png",
		"project5":     "/assets/image/rekammedis_app.png",
		"certifLogo2":  "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ2bF5xLMVYZuOW0z6KxrIeXZxnLvEDFgBAIA&s",
		"linkedin":     "https://www.linkedin.com/in/ahmadzkh/",
		"instagram":    "https://www.instagram.com/ahmadzkh_/",
		"twitter":      "https://x.com/ahmadzkh_/",
		"facebook":     "https://www.facebook.com/zaky.humami/",
		"extras1":      "/assets/image/img1.jpg",
		"extras2":      "/assets/image/img6.jpg",
		"extras3":      "/assets/image/img2.jpg",
		"extras4":      "/assets/image/img5.jpg",
		"extras5":      "/assets/image/img3.jpg",
		"extras6":      "/assets/image/img4.jpg",
	}

	// Execute main template with data
	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Oooppss.... error pages", http.StatusInternalServerError)
		return
	}
}
