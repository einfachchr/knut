package de.otto.obc.knut;

import org.springframework.http.HttpHeaders;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;

import java.util.HashMap;
import java.util.Map;

/**
 * Created by chr on 02.06.17.
 */
@Controller
public class RedController {

    @RequestMapping("/")
    public String index(@RequestHeader Map<String, String> headers, @RequestParam Map<String, String> reqParam, Model model) {

        HashMap<String, String> headerMap = new HashMap<>();
        HashMap<String, String> reqParamMap = new HashMap<>();

        headers.forEach((k, v) -> headerMap.put(k, v));
        reqParam.forEach((k, v) -> reqParamMap.put(k, v));

        model.addAttribute("headers", headerMap);
        model.addAttribute("requestParameter", reqParamMap);

        return "index";
    }
}
