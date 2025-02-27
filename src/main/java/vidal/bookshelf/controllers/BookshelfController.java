package vidal.bookshelf.controllers;

import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import vidal.bookshelf.repositories.BookRepository;

@RestController
public class BookshelfController {

    private final BookRepository bookRepository;

    public BookshelfController(BookRepository bookRepository) {
        this.bookRepository = bookRepository;
    }

    @RequestMapping("/index.html")
    public String getIndex(ModelMap model) {

        var books = bookRepository.findAll();

        System.out.println(books.size());

        model.addAttribute("books", books);

        return "index.html";
    }

}
